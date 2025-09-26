package hashing

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/cespare/xxhash/v2"
)

func TestReadFrom_1(t *testing.T) {
	tests := []struct {
		name    string
		r       io.Reader
		want    int64
		wantErr bool
	}{
		{
			name:    "test case 1",
			r:       strings.NewReader("hello world"),
			want:    11,
			wantErr: false,
		},
		{
			name:    "test case 2",
			r:       strings.NewReader(""),
			want:    0,
			wantErr: false,
		},
		{
			name:    "test case 3",
			r:       strings.NewReader("1234567890"),
			want:    10,
			wantErr: false,
		},
		{
			name:    "test case 4",
			r:       strings.NewReader("abcdefghijklmnopqrstuvwxyz"),
			want:    26,
			wantErr: false,
		},
		{
			name:    "test case 5",
			r:       strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
			want:    26,
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

			x := &xxhashReadFrom{
				buff:   make([]byte, 1024),
				Digest: xxhash.New(),
			}
			got, err := x.ReadFrom(tt.r)
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
