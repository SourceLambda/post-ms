package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// an alternative for a .env file
func SetEnvVars() {

	file, err := os.Open(".env")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		envVar := strings.Split(scanner.Text(), "=")
		if os.Getenv(envVar[0]) == "" {
			os.Setenv(envVar[0], envVar[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
