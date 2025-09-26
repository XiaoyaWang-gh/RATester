package limiter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_2(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		m    *manager
		args args
		want *item
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

			if got := tt.m.get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("manager.get() = %v, want %v", got, tt.want)
			}
		})
	}
}
