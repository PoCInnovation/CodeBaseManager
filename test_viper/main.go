package main

import (
     "fmt"
     "github.com/BurntSushi/toml"
     "os"
     "log"
)

type RepoTemplate struct {
     Language string `toml:"Language"`
     Sources struct{
          Modules []string
          Tests []string
     } `toml:"Sources"`
}

func printTomlInfo(filepath string, needed []string) {
     var v = make(map[string]RepoTemplate, 0)

     _, err := toml.DecodeFile(filepath, &v)
     if err != nil {
          fmt.Fprintln(os.Stderr, err)
          log.Fatal()
     }
     for _, toPrint := range needed {
          fmt.Println(v[toPrint])
     }
}

func main() {
     toFind := []string{"Sources", "Language"}

     for _, filepath := range os.Args[1:] {
          printTomlInfo(filepath, toFind)
     }
}