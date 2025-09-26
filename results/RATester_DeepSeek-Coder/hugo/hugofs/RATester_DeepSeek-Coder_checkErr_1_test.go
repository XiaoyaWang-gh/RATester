package hugofs

import (
	"fmt"
	"testing"
)

func TestCheckErr_1(t *testing.T) {
	type args struct {
		filename string
		err      error
	}
	tests := []struct {
		name string
		w    *Walkway
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

			if got := tt.w.checkErr(tt.args.filename, tt.args.err); got != tt.want {
				t.Errorf("Walkway.checkErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
