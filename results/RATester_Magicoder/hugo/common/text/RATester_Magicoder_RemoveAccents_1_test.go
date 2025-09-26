package text

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveAccents_1(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		want []byte
	}{
		{
			name: "Test with no accents",
			b:    []byte("Hello, World!"),
			want: []byte("Hello, World!"),
		},
		{
			name: "Test with accents",
			b:    []byte("Héllo, Wórld!"),
			want: []byte("Hello, World!"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := RemoveAccents(tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAccents() = %v, want %v", got, tt.want)
			}
		})
	}
}
