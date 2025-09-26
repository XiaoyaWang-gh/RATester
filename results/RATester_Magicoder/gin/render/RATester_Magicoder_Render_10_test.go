package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRender_10(t *testing.T) {
	testCases := []struct {
		name    string
		data    any
		wantErr bool
	}{
		{
			name:    "valid data",
			data:    "valid data",
			wantErr: false,
		},
		{
			name:    "invalid data",
			data:    make(chan int),
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := httptest.NewRecorder()
			r := AsciiJSON{Data: tc.data}
			err := r.Render(w)
			if (err != nil) != tc.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
