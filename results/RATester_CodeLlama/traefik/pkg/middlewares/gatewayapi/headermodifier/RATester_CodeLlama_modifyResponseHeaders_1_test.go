package headermodifier

import (
	"fmt"
	"net/http"
	"testing"
)

func TestModifyResponseHeaders_1(t *testing.T) {
	type args struct {
		res *http.Response
	}
	tests := []struct {
		name    string
		r       *responseHeaderModifier
		args    args
		wantErr bool
	}{
		{
			name: "test modifyResponseHeaders",
			r: &responseHeaderModifier{
				set: map[string]string{
					"test": "test",
				},
			},
			args: args{
				res: &http.Response{
					Header: make(http.Header),
				},
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

			if err := tt.r.modifyResponseHeaders(tt.args.res); (err != nil) != tt.wantErr {
				t.Errorf("modifyResponseHeaders() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
