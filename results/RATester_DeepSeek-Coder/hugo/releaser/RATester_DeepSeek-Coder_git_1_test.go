package releaser

import (
	"fmt"
	"testing"
)

func TestGit_1(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{
			name: "test case 1",
			args: []string{"arg1", "arg2"},
			want: "expected output",
		},
		// add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := git(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("git() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("git() = %v, want %v", got, tt.want)
			}
		})
	}
}
