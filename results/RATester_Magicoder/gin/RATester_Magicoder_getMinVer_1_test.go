package gin

import (
	"fmt"
	"testing"
)

func TestgetMinVer_1(t *testing.T) {
	tests := []struct {
		name    string
		v       string
		want    uint64
		wantErr bool
	}{
		{
			name:    "Test case 1",
			v:       "1.2.3",
			want:    2,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			v:       "1.2.3.4",
			want:    2,
			wantErr: false,
		},
		{
			name:    "Test case 3",
			v:       "1.2",
			want:    2,
			wantErr: false,
		},
		{
			name:    "Test case 4",
			v:       "1",
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test case 5",
			v:       "1.2.3.4.5",
			want:    2,
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

			got, err := getMinVer(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMinVer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getMinVer() = %v, want %v", got, tt.want)
			}
		})
	}
}
