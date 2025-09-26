package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_7(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		r       IndentedJSON
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			r: IndentedJSON{
				Data: map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
			args: args{
				w: httptest.NewRecorder(),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			r: IndentedJSON{
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

			if err := tt.r.Render(tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("IndentedJSON.Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
