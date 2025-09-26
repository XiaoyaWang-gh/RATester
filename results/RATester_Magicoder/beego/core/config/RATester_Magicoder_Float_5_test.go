package config

import (
	"fmt"
	"testing"
)

func TestFloat_5(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				key: "test_key",
			},
			want:    1.23,
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				key: "test_key_2",
			},
			want:    4.56,
			wantErr: false,
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

			got, err := Float(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Float() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Float() = %v, want %v", got, tt.want)
			}
		})
	}
}
