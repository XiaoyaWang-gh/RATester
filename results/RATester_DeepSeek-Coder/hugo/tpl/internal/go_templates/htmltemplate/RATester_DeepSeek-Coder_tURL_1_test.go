package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTURL_1(t *testing.T) {
	type args struct {
		c context
		s []byte
	}
	tests := []struct {
		name  string
		args  args
		want  context
		want1 int
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

			got, got1 := tURL(tt.args.c, tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tURL() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("tURL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
