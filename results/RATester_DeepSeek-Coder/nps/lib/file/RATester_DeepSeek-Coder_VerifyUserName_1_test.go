package file

import (
	"fmt"
	"testing"
)

func TestVerifyUserName_1(t *testing.T) {
	type args struct {
		username string
		id       int
	}
	tests := []struct {
		name string
		s    *DbUtils
		args args
		want bool
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

			if got := tt.s.VerifyUserName(tt.args.username, tt.args.id); got != tt.want {
				t.Errorf("DbUtils.VerifyUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}
