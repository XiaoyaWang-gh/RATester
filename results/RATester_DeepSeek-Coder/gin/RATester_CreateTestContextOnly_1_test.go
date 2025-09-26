package gin

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateTestContextOnly_1(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *Engine
	}

	tests := []struct {
		name string
		args args
		want *Context
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

			if got := CreateTestContextOnly(tt.args.w, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTestContextOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}
