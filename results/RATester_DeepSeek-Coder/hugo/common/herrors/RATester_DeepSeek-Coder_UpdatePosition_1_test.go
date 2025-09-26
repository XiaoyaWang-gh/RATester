package herrors

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/text"
)

func TestUpdatePosition_1(t *testing.T) {
	type args struct {
		pos text.Position
	}
	tests := []struct {
		name string
		fe   *fileError
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

			if got := tt.fe.UpdatePosition(tt.args.pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileError.UpdatePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
