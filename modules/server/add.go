package server

import (
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
		log.Println(err)
		return
	}
	if resp.StatusCode != http.StatusCreated {
		log.Println("Post request failed:", resp.Body)
		return
	}
}
