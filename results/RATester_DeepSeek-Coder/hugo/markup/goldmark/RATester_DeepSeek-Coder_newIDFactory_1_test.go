package goldmark

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewIDFactory_1(t *testing.T) {
	type args struct {
		idType string
	}
	tests := []struct {
		name string
		args args
		want *idFactory
	}{
		{
			name: "Test Case 1",
			args: args{
				idType: "test",
			},
			want: &idFactory{
				vals:   make(map[string]struct{}),
				idType: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newIDFactory(tt.args.idType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newIDFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
