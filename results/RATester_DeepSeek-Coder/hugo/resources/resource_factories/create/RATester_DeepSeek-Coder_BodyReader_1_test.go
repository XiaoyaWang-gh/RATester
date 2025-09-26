package create

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBodyReader_1(t *testing.T) {
	t.Run("returns nil when Body is nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		o := fromRemoteOptions{}
		reader := o.BodyReader()
		if reader != nil {
			t.Errorf("Expected nil, got %v", reader)
		}
	})

	t.Run("returns a bytes.Buffer when Body is not nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		o := fromRemoteOptions{Body: []byte("test")}
		reader := o.BodyReader()
		if reader == nil {
			t.Errorf("Expected a bytes.Buffer, got nil")
		} else {
			buf, ok := reader.(*bytes.Buffer)
			if !ok {
				t.Errorf("Expected a bytes.Buffer, got %T", reader)
			} else {
				if buf.String() != "test" {
					t.Errorf("Expected 'test', got %s", buf.String())
				}
			}
		}
	})
}
