package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	givenPrefix := "prefix"
	givenData := "data"
	givenSecureJSON := SecureJSON{
		Prefix: givenPrefix,
		Data:   givenData,
	}

	// when
	w := httptest.NewRecorder()
	err := givenSecureJSON.Render(w)

	// then
	if err != nil {
		t.Errorf("Render() error = %v", err)
		return
	}
	if w.Code != http.StatusOK {
		t.Errorf("Render() code = %v, want %v", w.Code, http.StatusOK)
	}
	if w.Body.String() != givenPrefix+givenData {
		t.Errorf("Render() body = %v, want %v", w.Body.String(), givenPrefix+givenData)
	}
}
