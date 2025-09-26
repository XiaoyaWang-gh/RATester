package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetWriteStr_1(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []byte
	}{
		{
			name: "Test case 1",
			args: []string{"Hello", "World"},
			want: []byte("HelloWorld"),
		},
		{
			name: "Test case 2",
			args: []string{"Golang", "is", "awesome"},
			want: []byte("Golangisawesome"),
		},
		{
			name: "Test case 3",
			args: []string{"1234567890", "abcdefghijklmnopqrstuvwxyz"},
			want: []byte("1234567890abcdefghijklmnopqrstuvwxyz"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetWriteStr(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWriteStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
