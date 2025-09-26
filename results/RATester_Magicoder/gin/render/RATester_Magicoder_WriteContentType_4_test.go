package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_4(t *testing.T) {
	testCases := []struct {
		name   string
		secure SecureJSON
		want   string
	}{
		{
			name:   "Test Case 1",
			secure: SecureJSON{Prefix: "test", Data: "test data"},
			want:   "application/json; charset=utf-8",
		},
		{
			name:   "Test Case 2",
			secure: SecureJSON{Prefix: "test", Data: 123},
			want:   "application/json; charset=utf-8",
		},
		{
			name:   "Test Case 3",
			secure: SecureJSON{Prefix: "test", Data: nil},
			want:   "application/json; charset=utf-8",
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
			tc.secure.WriteContentType(w)
			if got := w.Header().Get("Content-Type"); got != tc.want {
				t.Errorf("WriteContentType() = %v, want %v", got, tc.want)
			}
		})
	}
}
