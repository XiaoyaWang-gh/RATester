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
		// TODO: Add test cases.
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
