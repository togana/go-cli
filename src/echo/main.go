package main

import (
	"flag"
	"fmt"
	"strings"
)

var omitNewline = flag.Bool("n", false, "Do not print the trailing newline character.")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.Parse()
	var s string = strings.Join(flag.Args(), Space)

	if !*omitNewline {
		s += Newline
	}
	fmt.Print(s)
}
