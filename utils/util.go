package utils

import (
	"os"
	"strconv"
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

func ToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
