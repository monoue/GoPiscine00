package main

import (
	"fmt"
	"os"
	"strconv"
)

func putStair(num int) {
	const symbol = '*'

	for i := 0; i < num; i++ {
		fmt.Printf("%c", symbol)
	}
	fmt.Printf("%c", '\n')
}

func putStairs(num int) {
	i := 1

	for total := 1; total <= num; total += i {
		putStair(i)
		i++
	}
}

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	putStairs(num)
}
