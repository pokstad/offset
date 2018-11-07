package offset_test

import (
	"os"
	"testing"

	"github.com/pokstad/offset/offset"
)

func TestCalcRune(t *testing.T) {
	for _, tt := range []struct {
		name   string // filepath to fixture
		line   uint   // line number
		col    uint   // column rune offset
		offset uint   // rune offset
		errB   bool   // should we expect an error?
	}{
		{
			name:   "testdata/test.go",
			line:   3,
			col:    20,
			offset: 38,
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
