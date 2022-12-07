package main

import "testing"

const input = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

var dirTree = buildDirectoryTree(splitLines(input))
var dirSizes = getDirSizes(dirTree)

func Test_solvePart1(t *testing.T) {
	got := solvePart1(dirSizes)

	want := 95437
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test_solvePart2(t *testing.T) {
	got := solvePart2(dirTree, dirSizes)

	want := 24933642
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
