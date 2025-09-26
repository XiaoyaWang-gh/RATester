package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_4(t *testing.T) {
	type args struct {
		m   map[string][]string
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]string
		want1 bool
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

			c := &Context{}
			got, got1 := c.get(tt.args.m, tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
