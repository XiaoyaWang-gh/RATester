package herrors

import (
	"fmt"
	"testing"
)

func TestIsFeatureNotAvailableError_1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				err: fmt.Errorf("feature not available: %w", &FeatureNotAvailableError{}),
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				err: fmt.Errorf("some other error"),
			},
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

			if got := IsFeatureNotAvailableError(tt.args.err); got != tt.want {
				t.Errorf("IsFeatureNotAvailableError() = %v, want %v", got, tt.want)
			}
		})
	}
}
