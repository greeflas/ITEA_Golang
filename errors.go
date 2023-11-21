package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type EnvVarParseError struct {
	value string
}

func NewEnvVarParseError(value string) *EnvVarParseError {
	return &EnvVarParseError{value: value}
}

func (e EnvVarParseError) Error() string {
	return "invalid value"
}

func main() {
	const envFilePath = "./.env"

	err := loadEnv(envFilePath)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Oops, looks like file %q does not exist. You need to create it.\n", envFilePath)

		return
	}

	var parseError *EnvVarParseError
	if errors.As(err, &parseError) {
		fmt.Printf("File parse error: cannot parse line: %q\n", parseError.value)

		return
	}

	if err != nil {
		fmt.Printf("Fatal error: %s\n", err)

		return
	}

	fmt.Println(os.Getenv("FOO"), os.Getenv("BAR"))
}

func loadEnv(envFilePath string) error {
	f, err := os.Open(envFilePath)
	if err != nil {
		return fmt.Errorf("loadEnv: cannot open file: %q: %w", envFilePath, err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		name, value, err := parseEnvVar(line)
		if err != nil {
			return fmt.Errorf("loadEnv: invalid line: %q: %w", line, err)
		}

		if err := setEnvVar(name, value); err != nil {
			return fmt.Errorf("loadEnv: cannot set env var: %q: %w", name, err)
		}
	}

	f.Close()

	return nil
}

func parseEnvVar(envVar string) (string, string, error) {
	parts := strings.Split(envVar, "=")
	if len(parts) < 2 {
		return "", "", NewEnvVarParseError(envVar)
	}

	name, value := parts[0], parts[1]

	return name, value, nil
}

func setEnvVar(name, value string) error {
	return os.Setenv(name, value)
}
