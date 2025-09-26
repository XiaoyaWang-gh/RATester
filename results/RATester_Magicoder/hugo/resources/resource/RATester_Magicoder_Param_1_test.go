package resource

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestParam_1(t *testing.T) {
	type args struct {
		r        ResourceParamsProvider
		fallback maps.Params
		key      any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
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

			got, err := Param(tt.args.r, tt.args.fallback, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Param() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Param() = %v, want %v", got, tt.want)
			}
		})
	}
}
