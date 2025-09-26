package resources

import (
	"fmt"
	"testing"
)

func TestsetSourceFilenameIsHash_1(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		l    *genericResource
		args args
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

			tt.l.setSourceFilenameIsHash(tt.args.b)
		})
	}
}
