package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_2(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		r       TOML
		args    args
		wantErr bool
	}{
		{
			name: "Test Render",
			r: TOML{
				Data: map[string]string{
					"key": "value",
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
