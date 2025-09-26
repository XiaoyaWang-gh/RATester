package cors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPreflightHeader_1(t *testing.T) {
	type args struct {
		origin   string
		rMethod  string
		rHeaders string
	}
	tests := []struct {
		name string
		o    *Options
		args args
		want map[string]string
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

			if got := tt.o.PreflightHeader(tt.args.origin, tt.args.rMethod, tt.args.rHeaders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreflightHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
