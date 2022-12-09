package main

import (
	"fmt"
	"strings"
)

type Direction string

const (
	UP    Direction = "U"
	DOWN  Direction = "D"
	LEFT  Direction = "L"
	RIGHT Direction = "R"
)

func parseDirection(str string) Direction {
	if Direction(str) == UP {
		return UP
	}

	if Direction(str) == DOWN {
		return DOWN
	}

	if Direction(str) == LEFT {
		return LEFT
	}

	if Direction(str) == RIGHT {
		return RIGHT
	}

	msg := fmt.Sprintf("Couldn't parse %v to Direction!", str)
	panic(msg)
}

type Move struct {
	direction Direction
	steps     int
}

func parseMoves(lines []string) []Move {
	moves := make([]Move, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, " ")
		movie := Move{
			direction: parseDirection(split[0]),
			steps:     toInt(split[1]),
		}
		moves = append(moves, movie)
	}
	return moves
}

type Position struct {
	x, y int
}

func (p *Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func (p *Position) move(d Direction) {
	switch d {
	case UP:
		p.y += 1
	case DOWN:
		p.y -= 1
	case RIGHT:
		p.x += 1
	case LEFT:
		p.x -= 1
	default:
		msg := fmt.Sprintf("Implementation for direction %v doesn't exist!", d)
		panic(msg)
	}
}

func (p *Position) moveCloserTo(to Position) {
	yDiff := to.y - p.y
	if yDiff != 0 {
		if yDiff > 0 {
			p.y += 1
		} else {
			p.y -= 1
		}
	}

	xDiff := to.x - p.x
	if xDiff != 0 {
		if xDiff > 0 {
			p.x += 1
		} else {
			p.x -= 1
		}
	}
}
