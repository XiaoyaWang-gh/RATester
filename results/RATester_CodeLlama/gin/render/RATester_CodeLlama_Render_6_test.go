package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := String{
		Format: "Hello, %s!",
		Data: []any{
			"world",
		},
	}
	w := httptest.NewRecorder()

	// when
	err := r.Render(w)

	// then
	if err != nil {
		t.Errorf("Render() error = %v", err)
		return
	}
	if w.Code != http.StatusOK {
		t.Errorf("Render() code = %v, want %v", w.Code, http.StatusOK)
	}
	if w.Body.String() != "Hello, world!" {
		t.Errorf("Render() body = %v, want %v", w.Body.String(), "Hello, world!")
	}
}
