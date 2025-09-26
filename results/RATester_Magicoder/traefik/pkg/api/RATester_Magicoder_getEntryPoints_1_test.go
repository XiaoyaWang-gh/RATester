package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestgetEntryPoints_1(t *testing.T) {
	tests := []struct {
		name string
		h    Handler
		want []entryPointRepresentation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rw := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/entrypoints", nil)
			tt.h.getEntryPoints(rw, request)
			got := rw.Result()
			defer got.Body.Close()
			var gotEntryPoints []entryPointRepresentation
			err := json.NewDecoder(got.Body).Decode(&gotEntryPoints)
			if err != nil {
				t.Errorf("TestgetEntryPoints() error = %v", err)
				return
			}
			if !reflect.DeepEqual(gotEntryPoints, tt.want) {
				t.Errorf("TestgetEntryPoints() = %v, want %v", gotEntryPoints, tt.want)
			}
		})
	}
}
