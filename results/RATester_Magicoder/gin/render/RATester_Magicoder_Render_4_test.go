package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_4(t *testing.T) {
	tests := []struct {
		name     string
		r        Redirect
		wantErr  bool
		panicErr string
	}{
		{
			name: "Test case 1",
			r: Redirect{
				Code:     http.StatusOK,
				Request:  &http.Request{},
				Location: "http://example.com",
			},
			wantErr:  false,
			panicErr: "",
		},
		{
			name: "Test case 2",
			r: Redirect{
				Code:     http.StatusBadRequest,
				Request:  &http.Request{},
				Location: "http://example.com",
			},
			wantErr:  true,
			panicErr: "Cannot redirect with status code 400",
		},
		{
			name: "Test case 3",
			r: Redirect{
				Code:     http.StatusCreated,
				Request:  &http.Request{},
				Location: "http://example.com",
			},
			wantErr:  false,
			panicErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); r != nil {
					if r != tt.panicErr {
						t.Errorf("Render() panic = %v, wantErr %v", r, tt.panicErr)
					}
				}
			}()

			if err := tt.r.Render(httptest.NewRecorder()); (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
