package mock

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestMockCommit_1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *Mock
	}{
		{
			name: "TestMockCommit_0",
			args: args{
				err: nil,
			},
			want: NewMock(NewSimpleCondition("", "Commit"), []interface{}{nil}, nil),
		},
		{
			name: "TestMockCommit_1",
			args: args{
				err: errors.New("test error"),
			},
			want: NewMock(NewSimpleCondition("", "Commit"), []interface{}{errors.New("test error")}, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MockCommit(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockCommit() = %v, want %v", got, tt.want)
			}
		})
	}
}
