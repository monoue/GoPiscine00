package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG string = "Arguments is invalid."

func argsAreValid(args []string) bool {
	const size = 3

	if len(args) != size {
		return false
	}
	for i := 1; i < size; i++ {
		n, err := strconv.Atoi(args[i])
		if err != nil || (i == 2 && n == 0) {
			return false
		}
	}
	return true
}

func setResult(x, y int) string {
	sum := x + y
	difference := x - y
	product := x * y
	quotient := x / y

	return fmt.Sprintf(
		"sum: %v\ndifference: %v\nproduct: %v\nquotient: %v\n",
		sum,
		difference,
		product,
		quotient,
	)
}

func calculationStr(args []string) (string, bool) {
	if !argsAreValid(args) {
		return "", false
	}
	x, _ := strconv.Atoi(args[1])
	y, _ := strconv.Atoi(args[2])
	result := setResult(x, y)
	return result, true
}

func main() {
	s, ok := calculationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Print(s)
}
