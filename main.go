package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pokstad/offset/offset"
)

var (
	line = flag.Uint("line", 0, "line number in file")
	col  = flag.Uint("column", 0, "rune column offset in line")
	srcP = flag.String("source", "", "source file")
)

func main() {
	flag.Parse()

	var (
		src *os.File
		err error
	)

	// validate inputs

	switch {

	case *srcP != "":
		src, err = os.Open(*srcP)
		if err != nil {
			log.Fatalf("unable to open source file: %s", err)
		}

	default:
		// use STDIN when no file is specified
		src = os.Stdin
	}

	defer src.Close()

	runeOffset, err := offset.CalcRune(src, *line, *col)
	if err != nil {
		log.Fatalf("unable to calculate rune offset: %s", err)
	}

	fmt.Fprintf(os.Stdout, "%d\n", runeOffset)
}
