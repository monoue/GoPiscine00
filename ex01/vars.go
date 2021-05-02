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
	str := "42"
	putVarExplanation(str, fmt.Sprintf("%T", str))
	uintN := uint(42)
	putVarExplanation(fmt.Sprintf("%d", uintN), fmt.Sprintf("%T", uintN))
	intN := 42
	putVarExplanation(fmt.Sprintf("%d", intN), fmt.Sprintf("%T", intN))
	uint8N := uint8(42)
	putVarExplanation(fmt.Sprintf("%d", uint8N), fmt.Sprintf("%T", uint8N))
	int16N := int16(42)
	putVarExplanation(fmt.Sprintf("%d", int16N), fmt.Sprintf("%T", int16N))
	uint32N := uint32(42)
	putVarExplanation(fmt.Sprintf("%d", uint32N), fmt.Sprintf("%T", uint32N))
	int64N := int64(42)
	putVarExplanation(fmt.Sprintf("%d", int64N), fmt.Sprintf("%T", int64N))
	boolean := false
	putVarExplanation(fmt.Sprintf("%t", boolean), fmt.Sprintf("%T", boolean))
	float32N := float32(42)
	putVarExplanation(fmt.Sprintf("%g", float32N), fmt.Sprintf("%T", float32N))
	float64N := float64(42)
	putVarExplanation(fmt.Sprintf("%g", float64N), fmt.Sprintf("%T", float64N))
	// fmt.Printf("%v is %s %T.\n", str, aOrAn(str), str)
	// uin := uint32(42)
	// in, ui8n, i16n, ui32n, i64n, l

	// aOrAn(str)
	// fmt.Printf("%v is %s %T.\n", str, "a", str)
	// fmt.Printf("% is %s %T.\n", str, "a", str)
	// fmt.Println("42")
}
