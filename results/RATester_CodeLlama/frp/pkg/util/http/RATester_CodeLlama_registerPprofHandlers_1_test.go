package http

import (
	"fmt"
	"testing"

	"github.com/gorilla/mux"
)

func TestRegisterPprofHandlers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Server{}
	s.router = mux.NewRouter()
	s.registerPprofHandlers()
	// TODO: check if the pprof handlers are registered
}
