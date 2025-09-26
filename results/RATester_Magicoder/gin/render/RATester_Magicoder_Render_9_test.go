package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRender_9(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		r       Reader
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			r: Reader{
				ContentType:   "text/plain",
				ContentLength: 10,
				Reader:        strings.NewReader("Hello, World!"),
				Headers: map[string]string{
					"Content-Type": "text/plain",
				},
			},
			args: args{
				w: httptest.NewRecorder(),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			r: Reader{
				ContentType:   "text/plain",
				ContentLength: -1,
				Reader:        strings.NewReader("Hello, World!"),
				Headers: map[string]string{
					"Content-Type": "text/plain",
				},
			},
			args: args{
				w: httptest.NewRecorder(),
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			r: Reader{
				ContentType:   "text/plain",
				ContentLength: 10,
				Reader:        strings.NewReader(""),
				Headers: map[string]string{
					"Content-Type": "text/plain",
				},
			},
			args: args{
				w: httptest.NewRecorder(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.r.Render(tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
