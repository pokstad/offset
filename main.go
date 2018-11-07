package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pokstad/offset/offset"
)

var (
	line = flag.Int("line", -1, "line number in file (count starts at 1)")
	col  = flag.Int("column", -1, "column offset in line (count starts at 0)")
	srcP = flag.String("source", "", "source file")
)

func main() {
	flag.Parse()

	var (
		src *os.File
		err error
	)

	// validate inputs
	if *line < 1 || *col < 0 {
		log.Fatal("invalid values for line and column")
	}

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
