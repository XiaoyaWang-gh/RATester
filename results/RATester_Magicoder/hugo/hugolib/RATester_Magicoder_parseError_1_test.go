package hugolib

import (
	"fmt"
	"testing"
)

func TestparseError_1(t *testing.T) {
	type args struct {
		err    error
		input  []byte
		offset int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
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

			p := &pageState{}
			if err := p.parseError(tt.args.err, tt.args.input, tt.args.offset); (err != nil) != tt.wantErr {
				t.Errorf("parseError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
