package paths

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDoParse_1(t *testing.T) {
	type args struct {
		component string
		s         string
		p         *Path
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

			got, err := tt.pp.doParse(tt.args.component, tt.args.s, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathParser.doParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathParser.doParse() = %v, want %v", got, tt.want)
			}
		})
	}
}
