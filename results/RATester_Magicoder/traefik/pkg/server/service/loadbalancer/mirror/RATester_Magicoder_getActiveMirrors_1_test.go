package mirror

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetActiveMirrors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mirror := &Mirroring{
		mirrorHandlers: []*mirrorHandler{
			{
				Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}),
				percent: 100,
			},
		},
	}

	mirrors := mirror.getActiveMirrors()
	if len(mirrors) != 1 {
		t.Errorf("Expected 1 mirror, got %d", len(mirrors))
	}
}
