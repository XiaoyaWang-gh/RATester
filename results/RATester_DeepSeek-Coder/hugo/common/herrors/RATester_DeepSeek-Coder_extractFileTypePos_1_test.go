package herrors

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/text"
)

func TestExtractFileTypePos_1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 text.Position
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

			got, got1 := extractFileTypePos(tt.args.err)
			if got != tt.want {
				t.Errorf("extractFileTypePos() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("extractFileTypePos() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
