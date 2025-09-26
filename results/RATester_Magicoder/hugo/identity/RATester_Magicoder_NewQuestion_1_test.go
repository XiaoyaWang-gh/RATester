package identity

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewQuestion_1(t *testing.T) {
	type args struct {
		id Identity
	}
	tests := []struct {
		name string
		args args
		want *Question[any]
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

			if got := NewQuestion[any](tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuestion() = %v, want %v", got, tt.want)
			}
		})
	}
}
