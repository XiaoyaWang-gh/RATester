package herrors

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/text"
)

func TestNewFileErrorFromPos_1(t *testing.T) {
	type args struct {
		err error
		pos text.Position
	}
	tests := []struct {
		name string
		args args
		want FileError
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

			if got := NewFileErrorFromPos(tt.args.err, tt.args.pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileErrorFromPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
