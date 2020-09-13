package funcTests

import "testing"

type TestConfigFT struct {
	name, cfgPath string
	exp           ConfigFT
	err           error
}

func (tCfg *TestConfigFT) isValid(got *ConfigFT, err error) bool {

}

func TestNewConfigFT(t *testing.T) {
	tests := []TestConfigFT{
		{
			"Basic test file",
			"tests/basic.toml",
			ConfigFT{},
			nil,
		},
	}
	for _, test := range tests {
		got, err := NewConfigFT(test.cfgPath)
		if !test.isValid(got, err) {
			t.Errorf("%s")
		}
	}
}
