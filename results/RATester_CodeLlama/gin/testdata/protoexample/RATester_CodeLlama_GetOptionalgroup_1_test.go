package protoexample

import (
	"fmt"
	reflect "reflect"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestGetOptionalgroup_1(t *testing.T) {
	tests := []struct {
		name string
		test *Test
		want *Test_OptionalGroup
	}{
		{
			name: "test_get_optionalgroup_nil",
			test: &Test{},
			want: nil,
		},
		{
			name: "test_get_optionalgroup_not_nil",
			test: &Test{
				Optionalgroup: &Test_OptionalGroup{
					RequiredField: proto.String("test"),
				},
			},
			want: &Test_OptionalGroup{
				RequiredField: proto.String("test"),
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

			if got := tt.test.GetOptionalgroup(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test.GetOptionalgroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
