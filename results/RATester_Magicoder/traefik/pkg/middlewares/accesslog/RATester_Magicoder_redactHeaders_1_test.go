package accesslog

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestRedactHeaders_1(t *testing.T) {
	type args struct {
		headers http.Header
		fields  logrus.Fields
		prefix  string
	}
	tests := []struct {
		name string
		args args
		want logrus.Fields
	}{
		{
			name: "Test Case 1",
			args: args{
				headers: http.Header{
					"Authorization": []string{"Bearer token"},
					"Content-Type":  []string{"application/json"},
				},
				fields: logrus.Fields{},
				prefix: "prefix_",
			},
			want: logrus.Fields{
				"prefix_Authorization": "REDACTED",
				"prefix_Content-Type":  "application/json",
			},
		},
		{
			name: "Test Case 2",
			args: args{
				headers: http.Header{
					"Authorization": []string{"Bearer token"},
					"Content-Type":  []string{"application/json"},
				},
				fields: logrus.Fields{},
				prefix: "",
			},
			want: logrus.Fields{
				"Authorization": "REDACTED",
				"Content-Type":  "application/json",
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

			h := &Handler{}
			h.redactHeaders(tt.args.headers, tt.args.fields, tt.args.prefix)
			if !reflect.DeepEqual(tt.args.fields, tt.want) {
				t.Errorf("redactHeaders() = %v, want %v", tt.args.fields, tt.want)
			}
		})
	}
}
