package server

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	OFF         = false
	ON          = true
	BackendDir  = ".cbm/backend/"
	EnvFileName = ".env"
	LocalUrl    = "http://127.0.0.1"
)

var State = OFF
var Port = ""

func GetNewRequest(urlSuffix string) string {
	return fmt.Sprintf("%s:%s/%s", LocalUrl, Port, urlSuffix)
}

func GetBackendGlobalDirectory() (string, bool) {
	home := os.Getenv("HOME")
	if home == "" {
		log.Println("cannot load HOME environnement variable.")
		return "", false
	}
	var backendDir string
	if strings.HasSuffix(home, "/") {
		backendDir = home + BackendDir
	} else {
		backendDir = home + "/" + BackendDir
	}
	return backendDir, true
}

func getServerPort() bool {
	backendDir, ok := GetBackendGlobalDirectory()
	if !ok {
		return false
	}
	portFile := backendDir + EnvFileName
	err := godotenv.Load(portFile)
	if err != nil {
		log.Println(err)
		return false
	}
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
	healthUrl := fmt.Sprintf("%s:%s/health", LocalUrl, Port)
	response, err := http.Get(healthUrl)
	if err != nil || response.StatusCode != http.StatusOK {
		log.Println(err)
		State = OFF
		return State
	}
	State = ON
	return State
}
