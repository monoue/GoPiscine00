package main

import (
	"flag"
	"fmt"
	"regexp"
)

func isValidEmailAddressFormat(arg string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	// pattern := regexp.MustCompile(`^[a-zA-Z0-9\-._]+@[a-zA-Z0-9\-._]+\.[a-zA-Z]+$`)
	return pattern.MatchString(arg)
}

func isValidEmailAddress(s string) bool {
	return len(s) < 257 && isValidEmailAddressFormat(s)
}

func putEmailAddressValidation(arg string) {
	insertion := ""
	if !isValidEmailAddress(arg) {
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
	}
	// for _, arg := range flag.Args() {
	for _, arg := range args {
		putEmailAddressValidation(arg)
	}
}
