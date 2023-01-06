package main

import (
	"adian.com/advent_of_code_2022/utils"
	"strings"
)

func parse(data []string) map[point]bool {
	edges := parseEdges(data)
	result := make(map[point]bool)

	for _, structure := range edges {
		rockStructure := buildRockStructure(structure)
		for _, p := range rockStructure {
			result[p] = true
		}
	}

	return result
}

func buildRockStructure(structureEdges []point) []point {
	var structure []point
	var last point

	for _, edge := range structureEdges {
		isNotEmpty := last != (point{})
		if isNotEmpty {
			structure = append(structure, edge)

			if last.x == edge.x {
				for _, y := range toRange(edge.y, last.y) {
					p := point{
						x: edge.x,
						y: y,
					}
					structure = append(structure, p)
				}
			} else {
				for _, x := range toRange(edge.x, last.x) {
					p := point{
						x: x,
						y: edge.y,
					}
					structure = append(structure, p)
				}
			}
		}
		last = edge
	}

	return structure
}

func toRange(x int, y int) []int {
	var min, max int
	if x > y {
		max = x
		min = y
	} else {
		max = y
		min = x
	}

	var res []int
	for i := min; i <= max; i++ {
		res = append(res, i)
	}
	return res
}

func parseEdges(data []string) [][]point {
	edges := make([][]point, len(data))

	for i, line := range data {
		rawPoints := strings.Split(line, " -> ")

		for _, rawPoint := range rawPoints {
			split := strings.Split(rawPoint, ",")
			p := point{
				x: utils.ToInt(split[0]),
				y: utils.ToInt(split[1]),
			}
			edges[i] = append(edges[i], p)
		}
	}

	return edges
}
