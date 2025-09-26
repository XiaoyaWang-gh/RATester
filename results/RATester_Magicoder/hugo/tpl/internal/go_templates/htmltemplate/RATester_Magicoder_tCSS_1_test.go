package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttCSS_1(t *testing.T) {
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

			got, got1 := tCSS(tt.args.c, tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tCSS() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("tCSS() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
