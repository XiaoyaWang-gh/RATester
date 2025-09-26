package herrors

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/text"
)

func TestUpdatePosition_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fe := &fileError{}
	pos := text.Position{}
	pos.Filename = "filename"
	pos.LineNumber = 1
	pos.ColumnNumber = 1
	fe.UpdatePosition(pos)
	if fe.Position().Filename != "filename" {
		t.Errorf("fe.Position().Filename = %s, want %s", fe.Position().Filename, "filename")
	}
	if fe.Position().LineNumber != 1 {
		t.Errorf("fe.Position().LineNumber = %d, want %d", fe.Position().LineNumber, 1)
	}
	if fe.Position().ColumnNumber != 1 {
		t.Errorf("fe.Position().ColumnNumber = %d, want %d", fe.Position().ColumnNumber, 1)
	}
}
