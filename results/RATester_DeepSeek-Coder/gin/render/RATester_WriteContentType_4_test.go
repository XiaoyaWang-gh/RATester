package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_4(t *testing.T) {
	testCases := []struct {
		name string
		r    SecureJSON
		want string
	}{
		{
			name: "Test case 1",
			r:    SecureJSON{Prefix: "prefix", Data: "data"},
			want: "prefix",
		},
		{
			name: "Test case 2",
			r:    SecureJSON{Prefix: "another_prefix", Data: "data"},
			want: "another_prefix",
		},
		{
			name: "Test case 3",
			r:    SecureJSON{Prefix: "prefix", Data: "another_data"},
			want: "prefix",
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
			tc.r.WriteContentType(w)
			result := w.Header().Get("Content-Type")
			if result != tc.want {
				t.Errorf("Expected '%s', got '%s'", tc.want, result)
			}
		})
	}
}
