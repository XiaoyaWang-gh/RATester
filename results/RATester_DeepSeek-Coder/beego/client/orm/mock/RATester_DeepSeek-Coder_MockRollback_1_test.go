package mock

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestMockRollback_1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *Mock
	}{
		{
			name: "Test case 1",
			args: args{
				err: errors.New("test error"),
			},
			want: &Mock{
				cond: NewSimpleCondition("", "Rollback"),
				resp: []interface{}{errors.New("test error")},
				cb:   nil,
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

			if got := MockRollback(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockRollback() = %v, want %v", got, tt.want)
			}
		})
	}
}
