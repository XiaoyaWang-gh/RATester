package commands

import (
	"fmt"
	"os"
	"testing"
)

func TestIsTestRun_1(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{
			name: "Test with HUGO_TESTRUN set",
			env:  "true",
			want: true,
		},
		{
			name: "Test with HUGO_TESTRUN not set",
			env:  "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			os.Setenv("HUGO_TESTRUN", tt.env)
			defer os.Unsetenv("HUGO_TESTRUN")

			r := &rootCommand{}
			got := r.IsTestRun()
			if got != tt.want {
				t.Errorf("IsTestRun() = %v, want %v", got, tt.want)
			}
		})
	}
}
