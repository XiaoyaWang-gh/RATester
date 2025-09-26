package env

import (
	"fmt"
	"testing"
)

func TestGetRuntimeEnv_1(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		want    string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			key:     "key1",
			want:    "value1",
			wantErr: false,
		},
		{
			name:    "Test case 2",
			key:     "key2",
			want:    "value2",
			wantErr: false,
		},
		{
			name:    "Test case 3",
			key:     "key3",
			want:    "",
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

			got, err := GetRuntimeEnv(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRuntimeEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRuntimeEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
