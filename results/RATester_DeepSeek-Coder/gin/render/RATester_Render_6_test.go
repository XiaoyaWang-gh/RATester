package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRender_6(t *testing.T) {
	testCases := []struct {
		name    string
		s       String
		wantErr bool
	}{
		{
			name: "success",
			s: String{
				Format: "test",
				Data:   []any{"test"},
			},
			wantErr: false,
		},
		{
			name: "failure",
			s: String{
				Format: "test",
				Data:   []any{make(chan int)},
			},
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
			err := tc.s.Render(w)
			if (err != nil) != tc.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
