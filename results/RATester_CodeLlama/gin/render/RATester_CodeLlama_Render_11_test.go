package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := JsonpJSON{
		Callback: "callback",
		Data:     "data",
	}

	// when
	w := httptest.NewRecorder()
	err := r.Render(w)

	// then
	if err != nil {
		t.Errorf("Render() error = %v", err)
		return
	}

	if w.Code != http.StatusOK {
		t.Errorf("Render() code = %v, want %v", w.Code, http.StatusOK)
	}

	if w.Body.String() != "callback(data);" {
		t.Errorf("Render() body = %v, want %v", w.Body.String(), "callback(data);")
	}
}
