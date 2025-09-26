package herrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/text"
)

func TestNewFileErrorFromPos_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	err := errors.New("test error")
	pos := text.Position{Filename: "test.txt", LineNumber: 1, ColumnNumber: 1}
	fileError := NewFileErrorFromPos(err, pos)
	if fileError.Error() != "test error" {
		t.Errorf("NewFileErrorFromPos() = %v, want %v", fileError.Error(), "test error")
	}
	if fileError.Position() != pos {
		t.Errorf("NewFileErrorFromPos() = %v, want %v", fileError.Position(), pos)
	}
}
