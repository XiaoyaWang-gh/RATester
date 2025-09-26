package web

import (
	"fmt"
	"testing"
)

func Testne_1(t *testing.T) {

	tests := []struct {
		name    string
		arg1    interface{}
		arg2    interface{}
		want    bool
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			arg1:    "test",
			arg2:    "test",
			want:    false,
			wantErr: false,
		},
		{
			name:    "Test Case 2",
			arg1:    123,
			arg2:    456,
			want:    true,
			wantErr: false,
		},
		{
			name:    "Test Case 3",
			arg1:    "test",
			arg2:    123,
			want:    true,
			wantErr: false,
		},
		{
			name:    "Test Case 4",
			arg1:    nil,
			arg2:    nil,
			want:    false,
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

			got, err := ne(tt.arg1, tt.arg2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ne() = %v, want %v", got, tt.want)
			}
		})
	}
}
