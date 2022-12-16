package main

import (
	"strconv"
	"strings"
)

func parseInput(lines []string) valve {
	valvesMap := make(map[string]*valve, len(lines))
	tempLeadTo := make(map[string][]string)

	for _, line := range lines {
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ";", "")
		split := strings.Split(line, " ")

		flowRatePart := strings.Split(split[4], "=")[1]

		v := valve{
			name:     split[1],
			flowRate: toInt(flowRatePart),
		}
		valvesMap[v.name] = &v
		tempLeadTo[v.name] = split[9:]
	}

	for _, val := range valvesMap {
		leadToKeys := tempLeadTo[val.name]
		for _, key := range leadToKeys {
			val.leadTo = append(val.leadTo, valvesMap[key])
		}
	}

	return *valvesMap["AA"]
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}
