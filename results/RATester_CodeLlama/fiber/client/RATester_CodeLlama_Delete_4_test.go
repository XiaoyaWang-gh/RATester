package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/valyala/fasthttp"
	"github.com/zeebo/assert"
)

func TestDelete_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given: A test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// and: The method to test
	r := &Request{
		ctx: context.Background(),
		client: &Client{
			fasthttp: &fasthttp.Client{},
		},
	}

	// when: We call the method
	res, err := r.Delete(ts.URL)

	// then: There should be no error
	assert.NoError(t, err)

	// and: The response should be the expected one
	assert.Equal(t, http.StatusOK, res.StatusCode())
}
