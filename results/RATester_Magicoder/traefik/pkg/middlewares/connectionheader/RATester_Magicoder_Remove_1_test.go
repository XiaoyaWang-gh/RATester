package connectionheader

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRemove_1(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want *http.Request
	}{
		{
			name: "Test case 1",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"keep-alive"},
					"Upgrade":    []string{"websocket"},
				},
			},
			want: &http.Request{
				Header: http.Header{
					"Connection": []string{"keep-alive"},
					"Upgrade":    []string{"websocket"},
				},
			},
		},
		{
			name: "Test case 2",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"close"},
					"Upgrade":    []string{"websocket"},
				},
			},
			want: &http.Request{
				Header: http.Header{
					"Connection": []string{"close"},
					"Upgrade":    []string{"websocket"},
				},
			},
		},
		{
			name: "Test case 3",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"keep-alive"},
				},
			},
			want: &http.Request{
				Header: http.Header{
					"Connection": []string{"keep-alive"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := Remove(tt.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
