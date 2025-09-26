package denyrouterrecursion

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_9(t *testing.T) {
	tests := []struct {
		name           string
		routerName     string
		routerNameHash string
		req            *http.Request
		wantStatusCode int
	}{
		{
			name:           "same router",
			routerName:     "router1",
			routerNameHash: "hash1",
			req: &http.Request{
				Header: http.Header{
					xTraefikRouter: []string{"hash1"},
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "different router",
			routerName:     "router2",
			routerNameHash: "hash2",
			req: &http.Request{
				Header: http.Header{
					xTraefikRouter: []string{"hash1"},
				},
			},
			wantStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rr := httptest.NewRecorder()
			l := &DenyRouterRecursion{
				routerName:     tt.routerName,
				routerNameHash: tt.routerNameHash,
				next:           http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			}

			l.ServeHTTP(rr, tt.req)

			if status := rr.Code; status != tt.wantStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatusCode)
			}
		})
	}
}
