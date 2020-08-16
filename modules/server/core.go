package server

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	OFF = false
	ON  = true
)

var State = OFF
var Port = ""

func getServerPort() bool {
	home := os.Getenv("HOME")
	if home == "" {
		log.Println("cannot load HOME environnement variable.")
		return false
	}
	var portFile string
	if strings.HasSuffix(home, "/") {
		portFile = home + ".cbm/backend/.env"
	} else {
		portFile = home + "/.cbm/backend/.env"
	}
	err := godotenv.Load(portFile)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println(portFile)
	Port = os.Getenv("CBM_PORT")
	if Port == "" {
		log.Println("cannot load CBM_PORT environnement variable.")
		return false
	}
	return true
}

func GetServerState() bool {
	if Port == "" {
		if !getServerPort() {
			return false
		}
	}
	_, err := http.Get("http://127.0.0.1:8080/join")
	if err != nil {
		log.Println(err)
		State = OFF
	} else {
		State = ON
	}
	return State
}
