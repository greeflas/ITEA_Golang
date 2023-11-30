package config

import (
	"fmt"
	"os"
)

type Config struct {
	GitHubAccessToken string
}

func New() (*Config, error) {
	envVarName := "GITHUB_ACCESS_TOKEN"
	gitHubAccessToken, ok := os.LookupEnv(envVarName)
	if !ok {
		return nil, fmt.Errorf("required %s env var not found", envVarName)
	}

	return &Config{
		GitHubAccessToken: gitHubAccessToken,
	}, nil
}
