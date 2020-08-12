package codebase_test

import (
	"errors"
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/modules/codebase"
	"testing"
)

type TestConfigFT struct {
	name, cfgPath string
	exp           codebase.RepoTemplate
	err           error
}

func isContentOk(toPrint string, got []string, exp []string) error {
	if len(got) != len(exp) {
		return errors.New(fmt.Sprintf("Wrong %s len : expected (%d) <=> got (%d)", toPrint, len(exp), len(got)))
	}
	for i := 0; i != len(got); i += 1 {
		if got[i] != exp[i] {
			return errors.New(fmt.Sprintf("Wrong %s : expected (%s) <=> got (%s)", toPrint, exp[i], got[i]))
		}
	}
	return nil
}

func (tCfg *TestConfigFT) isValid(got *codebase.RepoTemplate, err error) error {
	if err != tCfg.err {
		return errors.New(fmt.Sprintf("Wrong error : got (%s) <=> expected (%s)", err, tCfg.err))
	}
	newErr := isContentOk("module", got.Modules, tCfg.exp.Modules)
	if newErr != nil {
		return newErr
	}

	newErr = isContentOk("language", got.Language, tCfg.exp.Language)
	if newErr != nil {
		return newErr
	}

	newErr = isContentOk("test", got.Tests, tCfg.exp.Tests)
	if newErr != nil {
		return newErr
	}

	return nil
}

func TestNewConfigFT(t *testing.T) {
	tests := []TestConfigFT{
		{
			"Basic test file",
			"test.toml",
			codebase.RepoTemplate{
				Language: []string{"C", "Go", "Python"},
				Modules:  []string{"parser", "repl", "cmd"},
				Tests:    []string{"tests", "olala", "did_you_see_my_skill?"},
			},
			nil,
		},
		{
			"File empty",
			"oops.toml",
			codebase.RepoTemplate{},
			nil,
		},
	}
	for _, test := range tests {
		got, err := codebase.GetTomlRepo(test.cfgPath)
		err = test.isValid(got, err)
		if err != nil {
			t.Errorf("Test %s failed : %s", test.name, err)
		}
	}
}
