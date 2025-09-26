package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_10(t *testing.T) {
	type args struct {
		config []Config
	}
	tests := []struct {
		name string
		args args
		want *App
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

			if got := New(tt.args.config...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
