package util

import (
	"os"
	"strings"
)

func ReadFileLines(filepath string) []string {
	raw, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	asString := string(raw)
	return strings.Split(asString, "\n")
}
