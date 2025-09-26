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

	srv := &Server{
		addr:   ":8080",
		router: mux.NewRouter(),
		authMiddleware: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			})
		},
	}

	called := false
	srv.RouteRegister(func(helper *RouterRegisterHelper) {
		assert.Equal(t, srv.router, helper.Router)
		assert.Equal(t, assets.FileSystem, helper.AssetsFS)
		assert.Equal(t, srv.authMiddleware, helper.AuthMiddleware)
		called = true
	})

	assert.True(t, called)
}
