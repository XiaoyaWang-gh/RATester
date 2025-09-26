package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteString_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	w := httptest.NewRecorder()
	format := "Hello, %s!"
	data := []any{"world"}
	// when
	err := WriteString(w, format, data)
	// then
	if err != nil {
		t.Errorf("WriteString() error = %v", err)
		return
	}
	if w.Code != http.StatusOK {
		t.Errorf("WriteString() code = %v, want %v", w.Code, http.StatusOK)
	}
	if w.Body.String() != "Hello, world!" {
		t.Errorf("WriteString() body = %v, want %v", w.Body.String(), "Hello, world!")
	}
}
