package pt

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func RedirectStdout() (*os.File, func()) {
	log.Println("Redirecting")
	saveStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return nil, nil
	}
	os.Stdout = w

	return r, func() {
		log.Println("Done redirecting")
		w.Close()
		os.Stdout = saveStdout
	}
}

func TestError(exp, got interface{}) error {
	if _, ok := exp.(string); ok {
		exp = strings.TrimSuffix(reflect.ValueOf(exp).String(), "\n")
		got = strings.TrimSuffix(reflect.ValueOf(got).String(), "\n")
	}
	return errors.New(fmt.Sprintf("\n--> Got: %v\n--> Exp: %v\n", got, exp))
}
