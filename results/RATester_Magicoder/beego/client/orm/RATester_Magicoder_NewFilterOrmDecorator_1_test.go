package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFilterOrmDecorator_1(t *testing.T) {
	type args struct {
		delegate     Ormer
		filterChains []FilterChain
	}
	tests := []struct {
		name string
		args args
		want Ormer
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

			if got := NewFilterOrmDecorator(tt.args.delegate, tt.args.filterChains...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterOrmDecorator() = %v, want %v", got, tt.want)
			}
		})
	}
}
