package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHeader_1(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		value  string
		header http.Header
		want   http.Header
	}{
		{
			name:  "Set header",
			key:   "Content-Type",
			value: "application/json",
			header: http.Header{
				"Content-Type": []string{"text/plain"},
			},
			want: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		{
			name:  "Delete header",
			key:   "Content-Type",
			value: "",
			header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			want: http.Header{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := &responseWriter{
				ResponseWriter: httptest.NewRecorder(),
			}
			c := &Context{
				Writer: w,
			}
			c.Header(tt.key, tt.value)
			if !reflect.DeepEqual(w.Header(), tt.want) {
				t.Errorf("Header() = %v, want %v", w.Header(), tt.want)
			}
		})
	}
}
