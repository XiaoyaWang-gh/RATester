package hugofs

import (
	"fmt"
	"os"
	"testing"
)

func TestIsWrite_1(t *testing.T) {
	type test struct {
		flag int
		want bool
	}

	tests := []test{
		{flag: os.O_RDONLY, want: false},
		{flag: os.O_WRONLY, want: true},
		{flag: os.O_RDWR, want: true},
		{flag: os.O_APPEND, want: false},
		{flag: os.O_CREATE, want: false},
		{flag: os.O_EXCL, want: false},
		{flag: os.O_SYNC, want: false},
		{flag: os.O_TRUNC, want: false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("isWrite(%d)", tt.flag), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := isWrite(tt.flag)
			if got != tt.want {
				t.Errorf("isWrite(%d) = %t, want %t", tt.flag, got, tt.want)
			}
		})
	}
}
