package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_15(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		r       XML
		args    args
		wantErr bool
	}{
		{
			name: "TestRender_Success",
			r: XML{
				Data: map[string]string{"key": "value"},
			},
			args: args{
				w: httptest.NewRecorder(),
			},
			wantErr: false,
		},
		{
			name: "TestRender_Fail",
			r: XML{
				Data: make(chan int),
			},
			args: args{
				w: httptest.NewRecorder(),
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

			err := tt.r.Render(tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("XML.Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
