package web

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestPayloadTooLarge_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	payloadTooLarge(rw, r)
	if rw.Code != 413 {
		t.Errorf("Expected 413, got %d", rw.Code)
	}
	if rw.Body.String() != `<br>The page you have requested is unavailable.
		 <br>Perhaps you are here because:<br><br>
		 <ul>
			<br>The request entity is larger than limits defined by server.
			<br>Please change the request entity and try again.
		 </ul>
		` {
		t.Errorf("Expected %s, got %s", `<br>The page you have requested is unavailable.
		 <br>Perhaps you are here because:<br><br>
		 <ul>
			<br>The request entity is larger than limits defined by server.
			<br>Please change the request entity and try again.
		 </ul>
		`, rw.Body.String())
	}
}
