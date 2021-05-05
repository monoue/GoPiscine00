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
