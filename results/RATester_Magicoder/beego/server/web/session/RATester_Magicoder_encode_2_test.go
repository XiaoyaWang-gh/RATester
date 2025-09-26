package session

import (
	"fmt"
	"reflect"
	"testing"
)

func Testencode_2(t *testing.T) {
	tests := []struct {
		name  string
		value []byte
		want  []byte
	}{
		{
			name:  "Test case 1",
			value: []byte("test"),
			want:  []byte("dGVzdA=="),
		},
		{
			name:  "Test case 2",
			value: []byte("hello"),
			want:  []byte("aGVsbG8="),
		},
		{
			name:  "Test case 3",
			value: []byte("world"),
			want:  []byte("d29ybGQ="),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := encode(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
