package create

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestBodyReader_1(t *testing.T) {
	tests := []struct {
		name string
		o    fromRemoteOptions
		want io.Reader
	}{
		{
			name: "nil body",
			o:    fromRemoteOptions{Body: nil},
			want: nil,
		},
		{
			name: "non-nil body",
			o:    fromRemoteOptions{Body: []byte("test")},
			want: bytes.NewBuffer([]byte("test")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.o.BodyReader(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BodyReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
