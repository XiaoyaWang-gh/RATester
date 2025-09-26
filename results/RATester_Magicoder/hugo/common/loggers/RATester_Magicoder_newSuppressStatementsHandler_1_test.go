package loggers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewSuppressStatementsHandler_1(t *testing.T) {
	type args struct {
		statements map[string]bool
	}
	tests := []struct {
		name string
		args args
		want *suppressStatementsHandler
	}{
		{
			name: "Test Case 1",
			args: args{
				statements: map[string]bool{"statement1": true, "statement2": false},
			},
			want: &suppressStatementsHandler{
				statements: map[string]bool{"statement1": true, "statement2": false},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				statements: map[string]bool{"statement3": true, "statement4": false},
			},
			want: &suppressStatementsHandler{
				statements: map[string]bool{"statement3": true, "statement4": false},
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

			if got := newSuppressStatementsHandler(tt.args.statements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSuppressStatementsHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
