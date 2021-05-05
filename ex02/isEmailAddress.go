package main

import (
	"flag"
	"fmt"
	"regexp"
)

var pattern = regexp.MustCompile(`^[a-zA-Z0-9\-._]+@[a-zA-Z0-9\-._]+\.[a-zA-Z]+$`)

func isValidEmailAddress(s string) bool {
	return len(s) < 257 && pattern.MatchString(s)
}

func putEmailAddressValidation(arg string) {
	var insertion string

	if isValidEmailAddress(arg) {
		insertion = ""
	} else {
		insertion = "NOT "
	}
	fmt.Printf("%v is %va valid email address.\n", arg, insertion)
}

func main() {
	const errMsg = "Invalid argument"

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println(errMsg)
		return
	}
	for _, arg := range args {
		putEmailAddressValidation(arg)
	}
}
