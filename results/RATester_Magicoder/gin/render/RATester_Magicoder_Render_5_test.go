package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_5(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r JSON
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				w: httptest.NewRecorder(),
				r: JSON{
					Data: "test data",
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				w: httptest.NewRecorder(),
				r: JSON{
					Data: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				w: httptest.NewRecorder(),
				r: JSON{
					Data: make(chan int),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.args.r.Render(tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
