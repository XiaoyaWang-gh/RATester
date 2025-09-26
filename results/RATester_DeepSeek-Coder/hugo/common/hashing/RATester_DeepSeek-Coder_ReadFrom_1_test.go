package hashing

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cespare/xxhash/v2"
)

func TestReadFrom_1(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    int64
		wantErr bool
	}{
		{
			name:    "Test case 1",
			data:    "test data",
			want:    8,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			data:    "",
			want:    0,
			wantErr: false,
		},
		{
			name:    "Test case 3",
			data:    "large data",
			want:    8,
			wantErr: false,
		},
		{
			name:    "Test case 4",
			data:    "error case",
			want:    5,
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

			r := strings.NewReader(tt.data)
			x := &xxhashReadFrom{
				buff:   make([]byte, 8),
				Digest: xxhash.New(),
			}
			got, err := x.ReadFrom(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadFrom() got = %v, want %v", got, tt.want)
			}
		})
	}
}
