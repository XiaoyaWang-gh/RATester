package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCopyBody_1(t *testing.T) {
	type args struct {
		MaxMemory int64
	}
	tests := []struct {
		name  string
		input *BeegoInput
		args  args
		want  []byte
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

			if got := tt.input.CopyBody(tt.args.MaxMemory); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BeegoInput.CopyBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
