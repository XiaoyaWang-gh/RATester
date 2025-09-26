package gin

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWrapF_1(t *testing.T) {
	type args struct {
		f http.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want HandlerFunc
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

			if got := WrapF(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapF() = %v, want %v", got, tt.want)
			}
		})
	}
}
