package binding

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDecodePlain_1(t *testing.T) {
	type testStruct struct {
		name string
		data []byte
		obj  any
		err  error
	}

	tests := []testStruct{
		{
			name: "Test with string type",
			data: []byte("test"),
			obj:  new(string),
			err:  nil,
		},
		{
			name: "Test with []byte type",
			data: []byte("test"),
			obj:  new([]byte),
			err:  nil,
		},
		{
			name: "Test with unknown type",
			data: []byte("test"),
			obj:  new(int),
			err:  fmt.Errorf("type (%T) unknown type", new(int)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := decodePlain(tt.data, tt.obj)
			if err != tt.err {
				t.Errorf("decodePlain() error = %v, wantErr %v", err, tt.err)
				return
			}

			if tt.obj != nil {
				switch v := tt.obj.(type) {
				case *string:
					if *v != string(tt.data) {
						t.Errorf("decodePlain() = %v, want %v", *v, string(tt.data))
					}
				case *[]byte:
					if !bytes.Equal(*v, tt.data) {
						t.Errorf("decodePlain() = %v, want %v", *v, tt.data)
					}
				}
			}
		})
	}
}
