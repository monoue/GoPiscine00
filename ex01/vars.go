package main

import "fmt"

func aOrAn(s string) string {
	if len(s) > 2 && (s[:3] == "int" || s == "*int") {
		return "an"
	}
	return "a"
}

func putVarExplanation(v, varType string) {
	fmt.Printf("%s is %s %s.\n", v, aOrAn(varType), varType)
}

func main() {
	type FortyTwo struct {
	}
	vars := []interface{}{
		"42",
		uint(42),
		42,
		uint8(42),
		int16(42),
		uint32(42),
		int64(42),
		false,
		float32(42),
		float64(42),
		(complex64)(42 + 0i),
		42 + 21i,
		FortyTwo{},
		[...]int{42},
		map[string]int{"42": 42},
		(*int)(nil),
		[]int{},
		chan int(nil),
		nil,
	}

	for _, v := range vars {
		if v == (*int)(nil) {
			putVarExplanation(fmt.Sprintf("%p", v), fmt.Sprintf("%T", v))
			continue
		}
		putVarExplanation(fmt.Sprintf("%v", v), fmt.Sprintf("%T", v))
	}
}
