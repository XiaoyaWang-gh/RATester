package herrors

import (
	"fmt"
	"testing"
)

func TestSetFilename_1(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		fe   *fileError
		args args
		want string
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

			if got := tt.fe.SetFilename(tt.args.filename); got.Position().Filename != tt.want {
				t.Errorf("fileError.SetFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
