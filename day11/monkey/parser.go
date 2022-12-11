package monkey

import (
	"strconv"
	"strings"
)

func Parse(lines []string) Monkeys {
	lines = normalizeLines(lines)

	var result Monkeys
	var currentMonkey *Monkey

	for _, line := range lines {
		split := splitLine(line)

		switch {
		case split[0] == "Monkey":
			id := parseMonkeyId(split[1])
			currentMonkey = &Monkey{
				Id: id,
			}
			result = append(result, currentMonkey)

		case split[0] == "Starting":
			for _, n := range split[2:] {
				currentMonkey.Items = append(currentMonkey.Items, toInt(n))
			}

		case split[0] == "Operation":
			currentMonkey.operation = parseOperation(split[3:])

		case split[0] == "Test":
			currentMonkey.ThrowTest = ThrowTest{
				DivisibleBy: toInt(getLast(split)),
			}

		case split[0] == "If" && split[1] == "true":
			currentMonkey.ThrowTest.ifTrue = parseMonkeyId(getLast(split))

		case split[0] == "If" && split[1] == "false":
			currentMonkey.ThrowTest.ifFalse = parseMonkeyId(getLast(split))
		}
	}

	return result
}

func splitLine(line string) []string {
	return strings.Split(line, " ")
}

func normalizeLines(lines []string) []string {
	var newLines []string
	for _, line := range lines {
		newLines = append(newLines, normalize(line))
	}

	return newLines
}

func normalize(line string) string {
	line = strings.ReplaceAll(line, ":", "")
	line = strings.ReplaceAll(line, ",", "")
	line = strings.TrimSpace(line)
	return line
}

func parseOperation(raw []string) Operation {
	return Operation{
		Num1:     raw[0],
		Operator: raw[1],
		Num2:     raw[2],
	}
}

func parseMonkeyId(str string) Id {
	return Id(toInt(str))
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)

	}
	return n
}

func getLast[T any](arr []T) T {
	return arr[len(arr)-1]
}
