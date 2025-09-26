package echo

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestNewResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}
	w := httptest.NewRecorder()
	r := NewResponse(w, e)

	if r.Writer != w {
		t.Errorf("Expected Writer to be %v, got %v", w, r.Writer)
	}
	if r.echo != e {
		t.Errorf("Expected echo to be %v, got %v", e, r.echo)
	}
	if r.Status != 0 {
		t.Errorf("Expected Status to be 0, got %v", r.Status)
	}
	if r.Size != 0 {
		t.Errorf("Expected Size to be 0, got %v", r.Size)
	}
	if r.Committed {
		t.Errorf("Expected Committed to be false, got %v", r.Committed)
	}
}
