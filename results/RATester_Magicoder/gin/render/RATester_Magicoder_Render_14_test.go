package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_14(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		r       PureJSON
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			r: PureJSON{
				Data: "test data",
			},
			args: args{
				w: httptest.NewRecorder(),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			r: PureJSON{
				Data: struct {
					Field string `json:"field"`
				}{
					Field: "test field",
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
