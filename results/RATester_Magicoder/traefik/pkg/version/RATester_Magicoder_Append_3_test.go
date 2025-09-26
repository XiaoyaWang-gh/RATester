package version

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestAppend_3(t *testing.T) {
	type args struct {
		router *mux.Router
	}
	tests := []struct {
		name string
		v    Handler
		args args
		want *mux.Router
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

			tt.v.Append(tt.args.router)
			if got := tt.args.router; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Append() = %v, want %v", got, tt.want)
			}
		})
	}
}
