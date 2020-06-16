package ft_types

import (
    "bufio"
    "errors"
    "fmt"
    "github.com/BurntSushi/toml"
    "github.com/PoCFrance/CodeBaseManager/modules/logs"
    "os"
    "strings"
)

type ftTestSuite struct {
    Vars map[string]string `toml:"vars"`
    Default ftDescription `toml:"default"`
    Tests []ftDescription `toml:"Test"`
}

type errConf struct {
    lineNumber int
    line string
}

func retrieveErr(toFind []string, line string, err *[]errConf, lineNumber int) {
    for _, separated := range toFind {
        if strings.Contains(line, separated){
            *err = append(*err, errConf{
                lineNumber: lineNumber,
                line: strings.TrimSpace(line)})
        }
    }
}

func getErrorToml(ignored []toml.Key, fp string) []errConf {
    f, _ := os.Open(fp)
    defer f.Close()
    err := make([]errConf, len(ignored))
    scanner := bufio.NewScanner(f)
    count := 0

    for scanner.Scan() {
        line := scanner.Text()
        for _, toFind := range ignored {
            toFind = strings.Split(toFind.String(), ".")
            retrieveErr(toFind, line, &err, count)
        }
        count += 1
    }
    return err
}

func NewTestSuite(cfgPath string) (*ftTestSuite, error) {
    cfg := &ftTestSuite{}

    md, err := toml.DecodeFile(cfgPath, cfg)
    if err != nil {
        return nil, err
    }
    // Checks if any key in the cfg were ignored
    ignored := md.Undecoded()
    if len(ignored) == 0 {
        return cfg, nil
    }
    ignoredErr := getErrorToml(ignored, cfgPath)
    logs.CBMLogs.Error("Syntax error at lines:")
    for _, k := range ignoredErr {
        if len(k.line) == 0 {
            continue
        }
        logs.CBMLogs.Error("%d:\t%s", k.lineNumber, k.line)
    }
    return nil, errors.New("Error while loading toml")
}

func (Test *ftDescription) Init(Default ftDescription) {
    Test.ftBasic.ApplyDefault(Default.ftBasic)
    Test.Expected.ApplyDefault(Default.Expected)
    Test.Interactions.ApplyDefault(Default.Interactions)
    Test.Options.ApplyDefault(Default.Options)
}

func (Test *ftDescription) PrintStateInteractions() {
    fmt.Println("Pre:")
    fmt.Println(Test.Interactions.Pre)
    fmt.Println("Post:")
    fmt.Println(Test.Interactions.Post)
    fmt.Println("Stdin file:")
    fmt.Println(Test.Interactions.StdinFile)
    fmt.Println("Stdin:")
    fmt.Println(Test.Interactions.Stdin)
    fmt.Println("StdinPipe:")
    fmt.Println(Test.Interactions.StdinPipe)
    fmt.Println("StdoutPipe:")
    fmt.Println(Test.Interactions.StdoutPipe)
    fmt.Println("StderrPipe:")
    fmt.Println(Test.Interactions.StderrPipe)
    fmt.Println("Env:")
    fmt.Println(Test.Interactions.Env)
    fmt.Println("AddEnv:")
    fmt.Println(Test.Interactions.AddEnv)
}

func (Test *ftDescription) PrintStateBasic() {
    fmt.Println("Name:")
    fmt.Println(Test.Name)
    fmt.Println("Desc:")
    fmt.Println(Test.Desc)
    fmt.Println("Bin:")
    fmt.Println(Test.ftBasic.Bin)
    fmt.Println("RefBin:")
    fmt.Println(Test.RefBin)
    fmt.Println("Args:")
    fmt.Println(Test.Args)
    fmt.Println("RefArgs:")
    fmt.Println(Test.RefArgs)
}

func (Test *ftDescription) PrintStateOptions() {
    fmt.Println("Time:")
    fmt.Println(Test.Options.Time)
    fmt.Println("Timeout")
    fmt.Println(Test.Options.Timeout)
    fmt.Println("ShouldFail:")
    fmt.Println(Test.Options.ShouldFail)
    fmt.Println("Repeat:")
    fmt.Println(Test.Options.Repeat)
}

func (Test *ftDescription) PrintStateExpected() {
    fmt.Println("Stdout:")
    fmt.Println(Test.Expected.Stdout)
    fmt.Println("Status:")
    fmt.Println(Test.Expected.Status)
    fmt.Println("Stderr:")
    fmt.Println(Test.Expected.Stderr)
    fmt.Println("StdoutFile")
    fmt.Println(Test.Expected.StdoutFile)
    fmt.Println("StderrFile")
    fmt.Println(Test.Expected.StderrFile)
}

func (cfg *ftTestSuite) BuildExec() error {
    for index, Test := range cfg.Tests {
        Test.Init(cfg.Default)
        cfg.Tests[index] = Test
    }
    /*for _, Test := range cfg.Tests {
        fmt.Println("-----------------------------------")
        Test.PrintStateBasic()
        Test.PrintStateExpected()
        Test.PrintStateInteractions()
        Test.PrintStateOptions()
    }*/
    // TODO: now that building is done, need to execute this
    return nil
}

