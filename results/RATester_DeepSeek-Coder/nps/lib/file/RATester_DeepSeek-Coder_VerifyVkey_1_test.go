package file

import (
	"fmt"
	"testing"
)

func TestVerifyVkey_1(t *testing.T) {
	type args struct {
		vkey string
		id   int
	}
	tests := []struct {
		name    string
		s       *DbUtils
		args    args
		wantRes bool
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

			if gotRes := tt.s.VerifyVkey(tt.args.vkey, tt.args.id); gotRes != tt.wantRes {
				t.Errorf("DbUtils.VerifyVkey() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
