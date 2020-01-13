package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Counter struct {
	lines int
	words int
	bytes int
}

type Flags struct {
	printLines bool
	printWords bool
	printBytes bool
}

func parseFlag() Flags {
	var flags Flags
	flag.BoolVar(&flags.printLines, "l", false, "The number of lines in each input file is written to the standard output.")
	flag.BoolVar(&flags.printWords, "w", false, "The number of words in each input file is written to the standard output.")
	flag.BoolVar(&flags.printBytes, "c", false, "The number of bytes in each input file is written to the standard output.")
	flag.Parse()

	if !flags.printLines && !flags.printWords && !flags.printBytes {
		flags.printLines = true
		flags.printWords = true
		flags.printBytes = true
	}

	return flags
}

func (c Counter) Print(filename string, flags Flags) {
	if flags.printLines {
		fmt.Printf(" %7d", c.lines)
	}
	if flags.printWords {
		fmt.Printf(" %7d", c.words)
	}
	if flags.printBytes {
		fmt.Printf(" %7d", c.bytes)
	}
	fmt.Printf(" %s\n", filename)
}

func (c *Counter) Count(file io.Reader) (bool, error) {
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return false, err
	}
	c.lines = bytes.Count(b, []byte{'\n'})
	c.bytes = len(b)
	c.words = len(bytes.Fields(b))
	return true, nil
}

func (c *Counter) Add(src Counter) {
	c.lines += src.lines
	c.bytes += src.bytes
	c.words += src.words
}

func main() {
	flags := parseFlag()
	var totalCount Counter

	filenames := flag.Args()

	if len(filenames) == 0 {
		var c Counter
		_, err := c.Count(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Print("", flags)
		return
	}

	for _, filename := range filenames {
		var c Counter
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer file.Close()

		_, err = c.Count(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		totalCount.Add(c)
		c.Print(filename, flags)
	}

	if len(filenames) > 1 {
		totalCount.Print("total", flags)
	}
}
