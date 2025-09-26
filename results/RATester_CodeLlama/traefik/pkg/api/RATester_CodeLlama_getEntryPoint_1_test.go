package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestGetEntryPoint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	entryPointID := "entryPointID"
	scapedEntryPointID := "entryPointID"
	entryPointID, err := url.PathUnescape(scapedEntryPointID)
	if err != nil {
		t.Errorf("unable to decode entryPointID %q: %s", scapedEntryPointID, err)
		return
	}
	rw := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/entrypoints/"+scapedEntryPointID, nil)
	if err != nil {
		t.Errorf("unable to create request: %s", err)
		return
	}
	h := Handler{
		staticConfig: static.Configuration{
			EntryPoints: map[string]*static.EntryPoint{
				entryPointID: {
					Address: "address",
				},
			},
		},
	}
	// when
	h.getEntryPoint(rw, request)
	// then
	if rw.Code != http.StatusOK {
		t.Errorf("unexpected status code: %d", rw.Code)
		return
	}
	var result entryPointRepresentation
	err = json.NewDecoder(rw.Body).Decode(&result)
	if err != nil {
		t.Errorf("unable to decode response: %s", err)
		return
	}
	if result.Name != entryPointID {
		t.Errorf("unexpected name: %s", result.Name)
		return
	}
	if result.EntryPoint.Address != "address" {
		t.Errorf("unexpected address: %s", result.EntryPoint.Address)
		return
	}
}
