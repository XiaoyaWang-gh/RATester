package text

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveAccents_1(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want []byte
	}{
		{
			name: "Test with accented characters",
			args: []byte("éèêëėę"),
			want: []byte("eeeeee"),
		},
		{
			name: "Test with non-accented characters",
			args: []byte("abcdefg"),
			want: []byte("abcdefg"),
		},
		{
			name: "Test with empty string",
			args: []byte(""),
			want: []byte(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := RemoveAccents(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAccents() = %v, want %v", got, tt.want)
			}
		})
	}
}
