package ft_types

import (
    "bufio"
    "errors"
    "fmt"
    "github.com/BurntSushi/toml"
    "github.com/PoCFrance/CodeBaseManager/modules/logs"
    "os"
    "reflect"
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

func searchInVars(toFind string, vars map[string]string) string {
    if !strings.Contains(toFind, "$") {
        return toFind
    }

    found := false
    for _, s := range reflect.ValueOf(vars).MapKeys() {
        str := s.String()
        if strings.Contains(toFind, str) {
            toFind = strings.Replace(toFind, fmt.Sprintf("$%s", str), vars[str], -1)
            found = true
        }
    }

    if found == false {
        logs.CBMLogs.Error(fmt.Sprintf("Variable not found in %s", toFind))
    }

    return toFind
}

func ApplyVarsString(toApplyString string, vars map[string]string) string {
    if toApplyString != "" {
        toApplyString = searchInVars(toApplyString, vars)
    }
    return toApplyString
}

func ApplyVarsTab(toApplyTab []string, vars map[string]string) []string {
    if len(toApplyTab) != 0 {
        for index, str := range toApplyTab {
            toApplyTab[index] = searchInVars(str, vars)
        }
    }
    return toApplyTab
}

func (Test *ftDescription) Init(Default ftDescription, vars map[string]string) {
    Test.ftBasic.ApplyDefault(Default.ftBasic, vars)
    Test.Expected.ApplyDefault(Default.Expected, vars)
    Test.Interactions.ApplyDefault(Default.Interactions, vars)
    Test.Options.ApplyDefault(Default.Options)
}

func (cfg *ftTestSuite) SetTestDefault() error {
    for index, Test := range cfg.Tests {
        Test.Init(cfg.Default, cfg.Vars)
        cfg.Tests[index] = Test
        if len(Test.Name) == 0 {
            return errors.New("no name found on the test")
        }
    }
    return nil
}

func (cfg *ftTestSuite) BuildExec() {
    if err := cfg.SetTestDefault(); err != nil {
        logs.CBMLogs.Error(err)
    }
    //Build done
    for _, Test := range cfg.Tests {
        for i := 0; i < Test.Options.Repeat; i++ { // Repeat handling
            //Pre command execution
            if err := QuickRun(Test.Interactions.Pre); err != nil {
                logs.CBMLogs.Error(err)
                continue
            }
            result, err := Test.Run()
            if err != nil {
                logs.CBMLogs.Error(err.Error())
                continue
            }
            result.AfterPipe(Test.Interactions.StdoutPipe, Test.Interactions.StderrPipe)
            //Post command execution
            if err := QuickRun(Test.Interactions.Post); err != nil {
                logs.CBMLogs.Error(err)
                continue
            }
            Test.GetResults(&result).Show(Test.Name)
            //TODO : check the result of the test
        }
    }
}

func (Test *ftDescription) GetResults(testOut *ftExecution) *ftResult {
    result := &ftResult{}
    if Test.RefBin != "" {
        Test.Bin = Test.RefBin
        Test.Args = Test.RefArgs
        ref, err := Test.Run()
        if err != nil {
            logs.CBMLogs.Error(err)
        }
        result.CompareToRef(&ref, testOut)
    } else {
        result.CompareToExp(testOut, &Test.Expected)
    }
    return result
}
