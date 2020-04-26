package funcTests

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

func Run(_ []string) {
	var cfg ConfigFT

	_, err := toml.DecodeFile(".cbm/template/ft.toml", &cfg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(cfg)
	}
}
