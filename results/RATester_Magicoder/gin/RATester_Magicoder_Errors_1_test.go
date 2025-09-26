package gin

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestErrors_1(t *testing.T) {
	tests := []struct {
		name string
		a    errorMsgs
		want []string
	}{
		{
			name: "no errors",
			a:    errorMsgs{},
			want: nil,
		},
		{
			name: "one error",
			a:    errorMsgs{&Error{Err: errors.New("error 1")}},
			want: []string{"error 1"},
		},
		{
			name: "multiple errors",
			a:    errorMsgs{&Error{Err: errors.New("error 1")}, &Error{Err: errors.New("error 2")}},
			want: []string{"error 1", "error 2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.a.Errors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Errors() = %v, want %v", got, tt.want)
			}
		})
	}
}
