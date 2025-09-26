package paths

import (
	"fmt"
	"testing"
)

func TestFileAndExt_1(t *testing.T) {
	type args struct {
		in string
		b  filepathPathBridge
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
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

			got, got1 := fileAndExt(tt.args.in, tt.args.b)
			if got != tt.want {
				t.Errorf("fileAndExt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("fileAndExt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
