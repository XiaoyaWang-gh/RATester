package web

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestFilter_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logFilter{}

	ctx0 := &beecontext.Context{
		Request: &http.Request{
			URL: &url.URL{
				Path: "/favicon.ico",
			},
		},
	}
	if l.Filter(ctx0) != true {
		t.Errorf("TestFilter failed, expected true, got false")
	}

	ctx1 := &beecontext.Context{
		Request: &http.Request{
			URL: &url.URL{
				Path: "/robots.txt",
			},
		},
	}
	if l.Filter(ctx1) != true {
		t.Errorf("TestFilter failed, expected true, got false")
	}

	ctx2 := &beecontext.Context{
		Request: &http.Request{
			URL: &url.URL{
				Path: "/static/file.txt",
			},
		},
	}
	if l.Filter(ctx2) != true {
		t.Errorf("TestFilter failed, expected true, got false")
	}

	ctx3 := &beecontext.Context{
		Request: &http.Request{
			URL: &url.URL{
				Path: "/notstatic/file.txt",
			},
		},
	}
	if l.Filter(ctx3) != false {
		t.Errorf("TestFilter failed, expected false, got true")
	}
}
