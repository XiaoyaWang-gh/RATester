package crypt

import (
	"fmt"
	"testing"
)

func TestMd5_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test case 1",
			arg:  "test",
			want: "098f6bcd4621d373cade4e832627b4f6",
		},
		{
			name: "Test case 2",
			arg:  "hello",
			want: "5d41402abc4b2a76b9719d911017c592",
		},
		{
			name: "Test case 3",
			arg:  "world",
			want: "7d793037a0760186574b0282f2f435e7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Md5(tt.arg); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}
