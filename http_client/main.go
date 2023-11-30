package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	goenv "github.com/greeflas/itea_goenv"
	"github.com/greeflas/itea_lessons/config"
	"github.com/greeflas/itea_lessons/dto"
	"io"
	"net/http"
	"time"
)

func main() {
	if err := goenv.LoadEnv("./.env"); err != nil {
		panic(err)
	}

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	printUserInfo()
	updateUserInfo(conf.GitHubAccessToken)
}

func printUserInfo() {
	resp, err := http.Get("https://api.github.com/users/greeflas")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	user := new(dto.GitHubUser)

	if err := json.Unmarshal(bodyBytes, user); err != nil {
		panic(err)
	}

	fmt.Printf("ID:\t%d\nName:\t%s\nBio:\t%s\n", user.Id, user.Name, user.Bio)
}

func updateUserInfo(githubAccessToken string) {
	const apiEndpoint = "https://api.github.com/user"

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	user := dto.GitHubUser{Bio: "I love Go!"}

	reqBody, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPatch, apiEndpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Authorization", "Bearer "+githubAccessToken)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("User was not updated: got %d status code\n", resp.StatusCode)
	}
}
