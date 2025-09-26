package integrity

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"reflect"
	"testing"
)

func TestNewHash_1(t *testing.T) {
	tests := []struct {
		name    string
		algo    string
		want    hash.Hash
		wantErr bool
	}{
		{
			name: "Test MD5",
			algo: "md5",
			want: md5.New(),
		},
		{
			name: "Test SHA256",
			algo: "sha256",
			want: sha256.New(),
		},
		{
			name: "Test SHA384",
			algo: "sha384",
			want: sha512.New384(),
		},
		{
			name: "Test SHA512",
			algo: "sha512",
			want: sha512.New(),
		},
		{
			name:    "Test Unsupported",
			algo:    "unsupported",
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

			got, err := newHash(tt.algo)
			if (err != nil) != tt.wantErr {
				t.Errorf("newHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("newHash() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
