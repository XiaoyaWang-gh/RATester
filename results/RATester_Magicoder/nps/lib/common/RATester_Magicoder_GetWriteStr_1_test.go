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
			args: []string{"test1", "test2"},
			want: []byte("test1test2"),
		},
		{
			name: "Test case 2",
			args: []string{"test3", "test4"},
			want: []byte("test3test4"),
		},
		// Add more test cases as needed
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
