package protoexample

import (
	"fmt"
	"testing"
)

func TestFile_test_proto_init_1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test file_test_proto_init",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			file_test_proto_init()
			if File_test_proto == nil {
				t.Errorf("file_test_proto_init() = %v, want %v", File_test_proto, "not nil")
			}
		})
	}
}
