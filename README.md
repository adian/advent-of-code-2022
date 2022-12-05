## Advent of Code 2022 in Golang
This project contains my solutions to [Advent of Code 2022](https://adventofcode.com/2022) . 
I wanted to learn Golang, and solving puzzles is the perfect opportunity to do it. 
That's why some solutions might not be optimal. I'm focusing here on new language than on algorithm's complexity.

## Usage
### Run solutions
Every puzzle is in a separated directory. For example, fifth day puzzle is in directory `day05`. 
Each directory with puzzle contains `main.go` file. Tu execute code, just run:
```shell
go run ./dayXX/
```
Where `dayXX` is the name of the directory, that contains puzzle.

### Run test

Run all tests
```shell
go test ./...
```

Run test of specific day
```shell
go test ./day03
```

