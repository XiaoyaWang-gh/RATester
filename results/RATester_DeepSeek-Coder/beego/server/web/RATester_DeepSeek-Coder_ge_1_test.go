package web

import (
	"fmt"
	"testing"
)

func TestGe_1(t *testing.T) {
	tests := []struct {
		name    string
		arg1    interface{}
		arg2    interface{}
		want    bool
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			arg1:    10,
			arg2:    20,
			want:    true,
			wantErr: false,
		},
		{
			name:    "Test Case 2",
			arg1:    20,
			arg2:    10,
			want:    false,
			wantErr: false,
		},
		{
			name:    "Test Case 3",
			arg1:    "10",
			arg2:    10,
			want:    false,
			wantErr: true,
		},
		{
			name:    "Test Case 4",
			arg1:    10,
			arg2:    "20",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Test Case 5",
			arg1:    "10",
			arg2:    "20",
			want:    true,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ge(tt.arg1, tt.arg2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ge() = %v, want %v", got, tt.want)
			}
		})
	}
}
