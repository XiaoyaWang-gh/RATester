package connectionheader

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRemoveConnectionHeaders_1(t *testing.T) {
	tests := []struct {
		name string
		h    http.Header
		want http.Header
	}{
		{
			name: "Test case 1",
			h:    http.Header{"Connection": []string{"close, keep-alive"}},
			want: http.Header{"Connection": []string{}},
		},
		{
			name: "Test case 2",
			h:    http.Header{"Connection": []string{"close, keep-alive"}, "Other-Header": []string{"value"}},
			want: http.Header{"Connection": []string{}, "Other-Header": []string{"value"}},
		},
		{
			name: "Test case 3",
			h:    http.Header{"Connection": []string{"close, keep-alive"}, "Other-Header": []string{"value1", "value2"}},
			want: http.Header{"Connection": []string{}, "Other-Header": []string{"value1", "value2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			removeConnectionHeaders(tt.h)
			if !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("removeConnectionHeaders() = %v, want %v", tt.h, tt.want)
			}
		})
	}
}
