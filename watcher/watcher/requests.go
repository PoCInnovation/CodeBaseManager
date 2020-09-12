package watcher

import (
    "errors"
    "fmt"
    "github.com/joho/godotenv"
    "github.com/PoCFrance/CodeBaseManager/ "
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

const (
    localUrl = "http://localhost:"
)

var port string

func initPort() {
    env, err := godotenv.Read(os.Getenv("HOME") + "/.cbm/backend/.env")
    if err != nil {
        log.Panicln("failed to load backend's port")
    }
    port = env["CBM_PORT"]
}

func getBackendPort() string {
    if port == "" {
        initPort()
    }
    return port
}

func GetProjectList() error {
    resp, err := http.Get(localUrl + getBackendPort() + "/project/list")
    if err != nil {
        return fmt.Errorf("failed to get projects list: %v", err)
    }
    if resp.StatusCode != http.StatusOK {
        return errors.New("no project found")
    }
    body, err := ioutil.ReadAll(resp.Body)
    repos := make()
    fmt.Println(string(body))
    return nil
}
