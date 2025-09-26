package ipwhitelist

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReject_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rw := httptest.NewRecorder()

	reject(ctx, rw)

	if rw.Code != http.StatusForbidden {
		t.Errorf("expected status code %d, got %d", http.StatusForbidden, rw.Code)
	}

	if rw.Body.String() != http.StatusText(http.StatusForbidden) {
		t.Errorf("expected body %q, got %q", http.StatusText(http.StatusForbidden), rw.Body.String())
	}
}
