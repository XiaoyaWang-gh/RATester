package hcontext

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewContextDispatcher_1(t *testing.T) {
	type testType struct {
		name string
		key  string
		want ContextDispatcher[testType]
	}

	tests := []testType{
		{
			name: "Test case 1",
			key:  "key1",
			want: NewContextDispatcher[testType, string]("key1"),
		},
		{
			name: "Test case 2",
			key:  "key2",
			want: NewContextDispatcher[testType, string]("key2"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewContextDispatcher[testType, string](tt.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContextDispatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
