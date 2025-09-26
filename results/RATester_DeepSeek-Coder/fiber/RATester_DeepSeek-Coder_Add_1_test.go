package fiber

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestAdd_1(t *testing.T) {
	app := New()
	registering := &Registering{
		app:  app,
		path: "/test",
	}

	handler := func(c Ctx) error {
		return c.SendString("Hello, World")
	}

	tests := []struct {
		name     string
		methods  []string
		expected string
	}{
		{
			name:     "GET request",
			methods:  []string{"GET"},
			expected: "Hello, World",
		},
		{
			name:     "POST request",
			methods:  []string{"POST"},
			expected: "Hello, World",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			registering.Add(test.methods, handler)

			req := httptest.NewRequest(test.methods[0], "/test", nil)
			resp, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			if string(body) != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, string(body))
			}
		})
	}
}
