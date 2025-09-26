package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestjoinRange_1(t *testing.T) {
	tests := []struct {
		name string
		c0   context
		rc   *rangeContext
		want context
	}{
		{
			name: "Test case 1",
			c0:   context{},
			rc: &rangeContext{
				breaks: []context{
					{
						state: stateError,
						err: &Error{
							Description: "Test error",
						},
					},
				},
			},
			want: context{
				state: stateError,
				err: &Error{
					Description: "at range loop break: Test error",
				},
			},
		},
		{
			name: "Test case 2",
			c0:   context{},
			rc: &rangeContext{
				continues: []context{
					{
						state: stateError,
						err: &Error{
							Description: "Test error",
						},
					},
				},
			},
			want: context{
				state: stateError,
				err: &Error{
					Description: "at range loop continue: Test error",
				},
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

			if got := joinRange(tt.c0, tt.rc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("joinRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
