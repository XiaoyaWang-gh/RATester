package protoexample

import (
	"fmt"
	reflect "reflect"
	"testing"
)

func TestFile_test_proto_rawDescGZIP_1(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{
			name: "Test case 1",
			want: []byte{0x12, 0x34, 0x56, 0x78},
		},
		{
			name: "Test case 2",
			want: []byte{0x01, 0x02, 0x03, 0x04},
		},
		{
			name: "Test case 3",
			want: []byte{0xff, 0xff, 0xff, 0xff},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := file_test_proto_rawDescGZIP()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("file_test_proto_rawDescGZIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
