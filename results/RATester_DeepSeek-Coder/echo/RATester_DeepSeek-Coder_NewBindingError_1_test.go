package echo

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewBindingError_1(t *testing.T) {
	type args struct {
		sourceParam   string
		values        []string
		message       interface{}
		internalError error
	}
	tests := []struct {
		name string
		args args
		want *BindingError
	}{
		{
			name: "Test case 1",
			args: args{
				sourceParam:   "field name",
				values:        []string{"value1", "value2"},
				message:       "Binding error",
				internalError: errors.New("internal error"),
			},
			want: &BindingError{
				Field: "field name",
				HTTPError: &HTTPError{
					Code:     http.StatusBadRequest,
					Message:  "Binding error",
					Internal: errors.New("internal error"),
				},
				Values: []string{"value1", "value2"},
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

			if got := NewBindingError(tt.args.sourceParam, tt.args.values, tt.args.message, tt.args.internalError); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBindingError() = %v, want %v", got, tt.want)
			}
		})
	}
}
