package ft_types

import (
    "bufio"
    "errors"
    "fmt"
    "github.com/BurntSushi/toml"
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
        fmt.Println(separated)
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
    NewLogger(ErrLog).Error("Syntax error at lines:")
    for _, k := range ignoredErr {
        if len(k.line) == 0 {
            continue
        }
        NewLogger(ErrLog).Error("%d:\t%s", k.lineNumber, k.line)
    }
    return nil, errors.New("Error while loading toml")
}

func (Test *ftDescription) Init(Default ftDescription) {
    Test.ftBasic.ApplyDefault(Default.ftBasic)
    Test.Expected.ApplyDefault(Default.Expected)
    Test.Interactions.ApplyDefault(Default.Interactions)
}

func (cfg *ftTestSuite) BuildExec() {
    for _, Test := range cfg.Tests {
        Test.Init(cfg.Default)
    }
}

