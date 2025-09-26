package rest

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateRouter_1(t *testing.T) {
	tests := []struct {
		name string
		p    *Provider
		want *mux.Router
	}{
		{
			name: "Test case 1",
			p: &Provider{
				Insecure: true,
			},
			want: func() *mux.Router {
				router := mux.NewRouter()
				router.Methods(http.MethodPut).Path("/api/providers/{provider}").Handler(&Provider{})
				return router
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.CreateRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.CreateRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
