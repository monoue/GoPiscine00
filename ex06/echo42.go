package main

import (
	"flag"
	"fmt"
)

func setFlags() (bool, string) {
	const defaultSep = " "
	noNewLine := flag.Bool("n", false, "omit trailing newline")
	sep := flag.String("s", defaultSep, "separator")
	flag.Parse()
	return *noNewLine, *sep
}

func echo42() {
	noNewLine, sep := setFlags()
	for i, arg := range flag.Args() {
		if i == 0 {
			fmt.Print(arg)
			continue
		}
		fmt.Print(sep)
		fmt.Print(arg)
	}
	if !noNewLine {
		fmt.Printf("%c", '\n')
	}
}

func main() {
	echo42()
}
