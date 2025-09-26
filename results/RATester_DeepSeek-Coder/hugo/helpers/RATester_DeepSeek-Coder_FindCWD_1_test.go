package helpers

import (
	"fmt"
	"os"
	"testing"
)

func TestFindCWD_1(t *testing.T) {
	type test struct {
		name        string
		want        string
		wantErr     bool
		setupEnv    func()
		teardownEnv func()
	}

	tests := []test{
		{
			name: "Test case 1",
			setupEnv: func() {
				os.Args = []string{"cmd", "/c", "echo", "Hello, World!"}
			},
			teardownEnv: func() {
				os.Args = []string{}
			},
			want:    ".",
			wantErr: false,
		},
		{
			name: "Test case 2",
			setupEnv: func() {
				os.Args = []string{"cmd", "/c", "echo", "Hello, World!"}
			},
			teardownEnv: func() {
				os.Args = []string{}
			},
			want:    ".",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.setupEnv()
			got, err := FindCWD()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindCWD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindCWD() = %v, want %v", got, tt.want)
			}
			tt.teardownEnv()
		})
	}
}
