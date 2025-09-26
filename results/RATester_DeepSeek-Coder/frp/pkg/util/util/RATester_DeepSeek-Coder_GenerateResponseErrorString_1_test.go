package util

import (
	"errors"
	"fmt"
	"testing"
)

func TestGenerateResponseErrorString_1(t *testing.T) {
	type args struct {
		summary  string
		err      error
		detailed bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				summary:  "Test summary",
				err:      errors.New("Test error"),
				detailed: true,
			},
			want: "Test error",
		},
		{
			name: "Test case 2",
			args: args{
				summary:  "Test summary",
				err:      errors.New("Test error"),
				detailed: false,
			},
			want: "Test summary",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GenerateResponseErrorString(tt.args.summary, tt.args.err, tt.args.detailed); got != tt.want {
				t.Errorf("GenerateResponseErrorString() = %v, want %v", got, tt.want)
			}
		})
	}
}
