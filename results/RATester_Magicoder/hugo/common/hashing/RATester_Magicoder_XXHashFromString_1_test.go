package hashing

import (
	"fmt"
	"testing"
)

func TestXXHashFromString_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				s: "test",
			},
			want:    10942973422275127737,
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				s: "another test",
			},
			want:    10942973422275127737,
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

			got, err := XXHashFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("XXHashFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("XXHashFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
