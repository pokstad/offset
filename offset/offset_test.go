package offset_test

import (
	"os"
	"testing"

	"github.com/pokstad/offset/offset"
)

func TestCalcRune(t *testing.T) {
	for _, tt := range []struct {
		name   string // filepath to fixture
		line   int    // line number
		col    int    // column rune offset
		offset int    // rune offset
		errB   bool   // should we expect an error?
	}{
		{
			name:   "testdata/test.go",
			line:   3,
			col:    20,
			offset: 37,
			errB:   false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			src, err := os.Open(tt.name)
			if err != nil {
				t.Fatalf("unable to open fixture: %s", err)
			}
			defer src.Close()

			o, err := offset.CalcRune(src, tt.line, tt.col)
			if !tt.errB && (err != nil) {
				t.Fatalf("unexpected calculation error: %s", err)
			}

			if o != tt.offset {
				t.Fatalf("offset does not match: %d vs %d", tt.offset, o)
			}
		})
	}
}
