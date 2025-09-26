package ports

import (
	"fmt"
	"testing"
)

func TestIsPortAvailable_1(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{port: 8080},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{port: 8081},
			want: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := (&Manager{}).isPortAvailable(tt.args.port); got != tt.want {
				t.Errorf("isPortAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}
