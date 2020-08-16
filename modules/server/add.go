package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Add(repoName string) {
	if !GetServerState() {
		log.Fatal("CodeBaseManager Backend not started")
	}

	actualPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	repoPath := filepath.Clean(actualPath + "/" + repoName)
	fmt.Println("Adding:", repoName)
	fmt.Println("Path:", repoPath)
	fmt.Println("[WIP]", GetServerState())
	addNewProject(repoName, repoPath)
}

func addNewProject(repoName, repoPath string) {
	url := GetApiUrl("project/add")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("projectName", repoName)
	q.Add("projectPath", repoPath)
	req.URL.RawQuery = q.Encode()

	resp, err := http.PostForm(req.URL.String(), q)
	if err != nil {
		log.Fatal(err)
	}
	//var res map[string]interface{}
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res)
	fmt.Println(res["name"])
	fmt.Println(res["path"])
}
