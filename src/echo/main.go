package main

import (
	"flag"
	"fmt"
)

var omitNewline = flag.Bool("n", false, "Do not print the trailing newline character.")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.Parse()

	var s string = ""

	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}

	if !*omitNewline {
		s += Newline
	}
	fmt.Print(s)
}
