package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_8(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		r       YAML
		args    args
		wantErr bool
	}{
		{
			name: "TestRender_Success",
			r: YAML{
				Data: map[string]interface{}{"key": "value"},
			},
			args:    args{w: httptest.NewRecorder()},
			wantErr: false,
		},
		{
			name: "TestRender_Fail",
			r: YAML{
				Data: make(chan int),
			},
			args:    args{w: httptest.NewRecorder()},
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
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
