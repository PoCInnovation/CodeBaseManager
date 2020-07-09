package repository

import (
    "bytes"
    "errors"
    "fmt"
    "github.com/logrusorgru/aurora"
    "log"
    "os"
    "os/exec"
    "regexp"
    "strconv"
    "strings"
)

type todoInfos struct {
    where int
    what string
}

func newTodoInfos(from []string) todoInfos {
    where, _ := strconv.Atoi(from[lineInFile])
    return todoInfos{
        where,
        from[todoTxt],
    }
}

type Todos map[string][]todoInfos

const (
    filepath = 1
    lineInFile = 2
    todoTxt = 3
    )

func (t Todos) PrintBasic() {
    for file, infos := range t {
        fmt.Println(aurora.Bold(file).Blue())
        for _, info := range infos {
            line := fmt.Sprintf("[%d]\t→", aurora.Bold(info.where).BrightBlue())
            fmt.Println(line, info.what)
        }
        fmt.Println("")
    }

}

func prettyPrint(cmd string) {
    execCmd := exec.Command("bash", "-c", cmd)
    execCmd.Stdout = os.Stdout
    _ = execCmd.Run()
}

func (t Todos) ctxWithCat(depth int) {
    for file, infos := range t {
        for _, info := range infos {
            fmt.Println(fmt.Sprintf("[%s:%d]", aurora.Bold(file).BrightBlue(), aurora.Bold(info.where)))
            head := info.where + depth
            tail := depth * 2
            cmd := fmt.Sprintf("cat %s | head -n %d | tail -n %d | cat", file, head, tail)
            prettyPrint(cmd)
        }
    }
}

func (t Todos) ctxWithBat(depth int) {
    var cmd string

    for file, infos := range t {
        for _, info := range infos {
            lower := info.where - depth
            if lower < 0 {
                lower = 0
            }
            upper := depth + info.where
            lang := strings.LastIndexByte(file, '.')
            cmd += fmt.Sprintf("bat %s -l %s -H %d -r %d:%d; ", file, file[lang+1:], info.where, lower, upper)
        }
    }
    prettyPrint(cmd)
}

func (t Todos) PrintWithCtx(depth int) {
    if _, err := exec.LookPath("bat"); err == nil {
        t.ctxWithBat(depth)
    } else {
        t.ctxWithCat(depth)
    }

}

func (t Todos) init(from []string) Todos {
    splitter := regexp.MustCompile("(.*):([0-9]*):.*(?i:TODO:(.*))")

    for _, todo := range from {
        if todo == "" {
            continue
        }
        res := splitter.FindStringSubmatch(todo)
        file := res[filepath]
        t[file] = append(t[file], newTodoInfos(res))
    }
    return t
}

func getFromCWD() (string, error) {
    var buf bytes.Buffer

    cmd := exec.Command("grep", "--color=auto", "--exclude-dir={.bzr,CVS,.git,.hg,.svn,.idea,.tox}", "-rni", "todo:")
    cmd.Stdout = &buf

    if err := cmd.Run(); err != nil {
        if err.Error() != "exit status 1" {
            return "", fmt.Errorf("couldn't retrieve todos: %v", err)
        } else {
            return "", errors.New("no todos found")
        }
    }
    return buf.String(), nil
}

func NewTodos() (Todos, error) {
    base, err := getFromCWD()
    if err != nil {
        return nil, fmt.Errorf("couldn't create todos: %v", err)
    }
    todoTab := strings.Split(base, "\n")
    newTodos := make(Todos)
    return newTodos.init(todoTab), nil
}

func DisplayTodos(ctxDepth int) {
    todos, err := NewTodos()
    if err != nil {
        log.Fatalln(err)
    }

    if ctxDepth == 0 {
        todos.PrintBasic()
    } else {
        todos.PrintWithCtx(ctxDepth)
    }
}