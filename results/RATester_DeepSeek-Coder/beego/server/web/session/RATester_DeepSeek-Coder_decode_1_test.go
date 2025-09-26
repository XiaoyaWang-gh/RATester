package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecode_1(t *testing.T) {
	tests := []struct {
		name    string
		value   []byte
		want    []byte
		wantErr bool
	}{
		{
			name:    "Test case 1",
			value:   []byte("dGVzdA=="),
			want:    []byte("test"),
			wantErr: false,
		},
		{
			name:    "Test case 2",
			value:   []byte("invalid"),
			want:    nil,
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

			got, err := decode(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
