package env

import (
	"bufio"
	"github.com/kongoole/minreuse-go/utils/log"
	"os"
	"strings"
)

func ParseEnv() {
	// read file line by line if exists
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal("failed to load .env file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		if text != "" && string(text[0]) != "#" {
			item := strings.Split(text, "=")
			os.Setenv(item[0], item[1])
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
