package livereload

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gorilla/websocket"
)

func TestReader_1(t *testing.T) {
	type fields struct {
		ws     *websocket.Conn
		send   chan []byte
		closer sync.Once
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &connection{
				ws:     tt.fields.ws,
				send:   tt.fields.send,
				closer: tt.fields.closer,
			}
			c.reader()
		})
	}
}
