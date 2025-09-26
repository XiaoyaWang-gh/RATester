package allconfig

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWrapFileError_1(t *testing.T) {
	type args struct {
		err      error
		filename string
	}
	tests := []struct {
		name string
		l    configLoader
		args args
		want error
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

			if got := tt.l.wrapFileError(tt.args.err, tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configLoader.wrapFileError() = %v, want %v", got, tt.want)
			}
		})
	}
}
