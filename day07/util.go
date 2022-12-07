package main

import (
	"strconv"
	"strings"
)

func splitLines(raw string) []string {
	return strings.Split(raw, "\n")
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
