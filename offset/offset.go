package offset

import (
	"bufio"
	"io"
)

// CalcRune will return the byte offset for the specified line and column of
// the reader (e.g. source code file). Both line and column indexes start at 1.
func CalcRune(r io.Reader, line, col uint) (uint, error) {
	var (
		s  = bufio.NewScanner(r)
		bc = uint(0)
	)

	for i := uint(1); s.Scan() && i < line; i++ {
		if err := s.Err(); err != nil {
			return 0, err
		}

		bc += uint(len(s.Text()) + 1) // add newline to text length
	}

	return bc + col, nil
}
