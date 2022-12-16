package utils

import (
	"os"
	"strings"
)

func ReadFileLines(filepath string) []string {
	asString := ReadFile(filepath)
	return strings.Split(asString, "\n")
}

func ReadFile(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	data := string(file)
	return data
}
