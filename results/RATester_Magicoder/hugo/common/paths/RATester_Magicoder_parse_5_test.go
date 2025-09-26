package paths

import (
	"fmt"
	"reflect"
	"testing"
)

func Testparse_5(t *testing.T) {
	type args struct {
		component string
		s         string
	}
	tests := []struct {
		name    string
		pp      *PathParser
		args    args
		want    *Path
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

			got, err := tt.pp.parse(tt.args.component, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathParser.parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathParser.parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
