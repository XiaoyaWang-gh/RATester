package web

import (
	"fmt"
	"testing"
)

func TestinitAddr_1(t *testing.T) {
	app := &HttpServer{
		Cfg: &Config{
			Listen: Listen{
				HTTPAddr: "",
			},
		},
	}

	tests := []struct {
		name     string
		addr     string
		expected string
	}{
		{"empty", "", ""},
		{"only_host", "localhost", "localhost"},
		{"only_port", ":8080", ""},
		{"host_and_port", "localhost:8080", "localhost"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			app.initAddr(test.addr)
			if app.Cfg.Listen.HTTPAddr != test.expected {
				t.Errorf("Expected HTTPAddr to be %s, but got %s", test.expected, app.Cfg.Listen.HTTPAddr)
			}
		})
	}
}
