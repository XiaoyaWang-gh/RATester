package create

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestResponseToData_1(t *testing.T) {
	type args struct {
		res      *http.Response
		readBody bool
	}
	tests := []struct {
		name string
		args args
		want map[string]any
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

			if got := responseToData(tt.args.res, tt.args.readBody); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("responseToData() = %v, want %v", got, tt.want)
			}
		})
	}
}
