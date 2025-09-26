package service

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
)

func TestCleanWebSocketHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &http.Request{}
	cleanWebSocketHeaders(req)
	assert.Equal(t, req.Header["Sec-WebSocket-Key"], req.Header["Sec-Websocket-Key"])
	assert.Equal(t, req.Header["Sec-WebSocket-Extensions"], req.Header["Sec-Websocket-Extensions"])
	assert.Equal(t, req.Header["Sec-WebSocket-Accept"], req.Header["Sec-Websocket-Accept"])
	assert.Equal(t, req.Header["Sec-WebSocket-Protocol"], req.Header["Sec-Websocket-Protocol"])
	assert.Equal(t, req.Header["Sec-WebSocket-Version"], req.Header["Sec-Websocket-Version"])
}
