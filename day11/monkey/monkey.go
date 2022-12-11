package monkey

import "fmt"

type Id int

type Operation struct {
	Num1     string
	Num2     string
	Operator string
}

func (o *Operation) calculate(item int) int {
	toNumber := func(str string) int {
		if str == "old" {
			return item
		} else {
			return toInt(str)
		}
	}

	num1, num2 := toNumber(o.Num1), toNumber(o.Num2)

	switch o.Operator {
	case "*":
		return num1 * num2
	case "+":
		return num1 + num2
	default:
		msg := fmt.Sprintf("operator %v is not implemented!", o.Operator)
		panic(msg)
	}
}

type ThrowTest struct {
	DivisibleBy int
	ifTrue      Id
	ifFalse     Id
}

type Monkey struct {
	Id                  Id
	Items               []int
	operation           Operation
	ThrowTest           ThrowTest
	NumberOfInspections int
	WorryLevelModifier  func(worryLevel int) int
}

func (m *Monkey) ThrowFirstItem() (to Id, what int, ok bool) {
	if len(m.Items) == 0 {
		return 0, 0, false
	}
	m.NumberOfInspections += 1

	firstItem := m.Items[0]
	m.Items = m.Items[1:]

	worryLevel := m.calculateWorryLevel(firstItem)
	return m.whereThrow(worryLevel), worryLevel, true
}

func (m *Monkey) whereThrow(worryLevel int) Id {
	if worryLevel%m.ThrowTest.DivisibleBy == 0 {
		return m.ThrowTest.ifTrue
	}
	return m.ThrowTest.ifFalse
}

func (m *Monkey) calculateWorryLevel(i int) int {
	result := m.operation.calculate(i)
	if m.WorryLevelModifier != nil {
		result = m.WorryLevelModifier(result)
	}
	return result
}

type Monkeys []*Monkey
