package main

import (
	"flag"
	"os"
)

var omitNewline = flag.Bool("n", false, "don't print final newline")

const (
	Space   = ""
	Newline = "\n"
)

func main() {
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		s = ""
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
		if !*omitNewline {
			s += Newline
		}
		os.Stdout.WriteString(s)
	}
}
