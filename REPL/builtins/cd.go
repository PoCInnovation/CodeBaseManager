package builtins

import (
	"fmt"
	"os"
)

func updateEnv(newPwd string, oldPwd string) {
	err := os.Setenv("PWD", newPwd)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Setenv("OLDPWD", oldPwd)
	if err != nil {
		fmt.Println(err)
	}
}

func CD(av []string) {
	if len(av) != 2 {
		fmt.Fprintln(os.Stderr, "cd: requires exactly one argument")
		return
	}

	oldPwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	err = os.Chdir(av[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	newPwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	updateEnv(newPwd, oldPwd)
}
