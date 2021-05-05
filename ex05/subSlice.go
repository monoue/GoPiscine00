package main

import "fmt"

func subSlice(slice []int, begin int, length int, capacity int) []int {
	if begin < 0 || length < 0 {
		return nil
	}
	var capa int
	if capacity < length {
		capa = length
	} else {
		capa = capacity
	}
	subslice := make([]int, length, capa)
	origSize := len(slice)
	for i := 0; i < length && begin+i < origSize; i++ {
		subslice[i] = slice[begin+i]
	}
	return subslice
}

func main() {
	var orig = []int{0, 1, 2, 3, 4, 5}
	var ret []int
	ret = subSlice(orig, 0, 3, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
	ret = subSlice(orig, 2, 7, 10)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
	ret = subSlice(orig, 2, 7, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
}
