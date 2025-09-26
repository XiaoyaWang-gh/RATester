package herrors

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/text"
)

func TestExtractPosition_1(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name       string
		args       args
		wantPos    text.Position
		wantString string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotPos := extractPosition(tt.args.e)
			if gotPos.Filename != tt.wantPos.Filename {
				t.Errorf("extractPosition() = %v, want %v", gotPos.Filename, tt.wantPos.Filename)
			}
			if gotPos.Offset != tt.wantPos.Offset {
				t.Errorf("extractPosition() = %v, want %v", gotPos.Offset, tt.wantPos.Offset)
			}
			if gotPos.LineNumber != tt.wantPos.LineNumber {
				t.Errorf("extractPosition() = %v, want %v", gotPos.LineNumber, tt.wantPos.LineNumber)
			}
			if gotPos.ColumnNumber != tt.wantPos.ColumnNumber {
				t.Errorf("extractPosition() = %v, want %v", gotPos.ColumnNumber, tt.wantPos.ColumnNumber)
			}
			if gotString := gotPos.String(); gotString != tt.wantString {
				t.Errorf("extractPosition().String() = %v, want %v", gotString, tt.wantString)
			}
		})
	}
}
