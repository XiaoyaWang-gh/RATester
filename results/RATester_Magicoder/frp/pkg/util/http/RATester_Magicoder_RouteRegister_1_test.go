package http

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/fatedier/frp/assets"
	"github.com/gorilla/mux"
	"github.com/zeebo/assert"
)

func TestRouteRegister_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		addr:   "localhost:8080",
		router: mux.NewRouter(),
		authMiddleware: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Add authentication logic here
				next.ServeHTTP(w, r)
			})
		},
	}

	called := false
	server.RouteRegister(func(helper *RouterRegisterHelper) {
		called = true
		assert.Equal(t, server.router, helper.Router)
		assert.Equal(t, assets.FileSystem, helper.AssetsFS)
		assert.Equal(t, server.authMiddleware, helper.AuthMiddleware)
	})

	assert.True(t, called)
}
