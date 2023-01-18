package main

import (
	"strings"

	"adian.com/advent_of_code_2022/utils"
)

var lineNormalizer = strings.NewReplacer(
	":", "",
	",", "",
	"=", " ",
)

func parse(raw []string) []pair {
	var result []pair

	for _, line := range raw {
		line = lineNormalizer.Replace(line)
		split := strings.Split(line, " ")

		temp := pair{
			sensor: position{
				x: utils.ToInt(split[3]),
				y: utils.ToInt(split[5]),
			},
			beacon: position{
				x: utils.ToInt(split[11]),
				y: utils.ToInt(split[13]),
			},
		}
		result = append(result, temp)
	}

	return result
}
