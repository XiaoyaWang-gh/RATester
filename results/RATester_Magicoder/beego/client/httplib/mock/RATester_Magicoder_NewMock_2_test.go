package mock

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewMock_2(t *testing.T) {
	type args struct {
		con  RequestCondition
		resp *http.Response
		err  error
	}
	tests := []struct {
		name string
		args args
		want *Mock
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

			if got := NewMock(tt.args.con, tt.args.resp, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMock() = %v, want %v", got, tt.want)
			}
		})
	}
}
