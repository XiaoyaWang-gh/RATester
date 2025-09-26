package crypt

import (
	"fmt"
	"testing"
)

func TestMd5_1(t *testing.T) {
	t.Run("TestMd5", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		tests := []struct {
			name string
			arg  string
			want string
		}{
			{
				name: "Test case 1",
				arg:  "hello",
				want: "5d41402abc4b2a76b9719d911017c592",
			},
			{
				name: "Test case 2",
				arg:  "world",
				want: "7d793037a0760186574b0282f2f435e7",
			},
			{
				name: "Test case 3",
				arg:  "golang",
				want: "c20ad4d5685ac3703064c422299c5218",
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
	})
}
