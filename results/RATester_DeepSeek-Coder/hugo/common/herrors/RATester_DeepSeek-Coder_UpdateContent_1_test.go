package herrors

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestUpdateContent_1(t *testing.T) {
	type args struct {
		r           io.Reader
		linematcher LineMatcherFn
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

			if got := tt.fe.UpdateContent(tt.args.r, tt.args.linematcher); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileError.UpdateContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
