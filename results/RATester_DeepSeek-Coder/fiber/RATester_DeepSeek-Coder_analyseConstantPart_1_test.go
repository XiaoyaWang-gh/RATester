package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAnalyseConstantPart_1(t *testing.T) {
	type args struct {
		pattern           string
		nextParamPosition int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 *routeSegment
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

			rp := &routeParser{}
			got, got1 := rp.analyseConstantPart(tt.args.pattern, tt.args.nextParamPosition)
			if got != tt.want {
				t.Errorf("routeParser.analyseConstantPart() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("routeParser.analyseConstantPart() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
