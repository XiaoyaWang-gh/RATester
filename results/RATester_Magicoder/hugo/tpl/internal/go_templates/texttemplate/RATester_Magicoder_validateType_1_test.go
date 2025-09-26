package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestvalidateType_1(t *testing.T) {
	type testCase struct {
		name    string
		value   reflect.Value
		typ     reflect.Type
		want    reflect.Value
		wantErr bool
	}

	tests := []testCase{
		{
			name:    "valid value",
			value:   reflect.ValueOf(1),
			typ:     reflect.TypeOf(1),
			want:    reflect.ValueOf(1),
			wantErr: false,
		},
		{
			name:    "invalid value",
			value:   reflect.ValueOf(1),
			typ:     reflect.TypeOf("string"),
			want:    reflect.ValueOf(1),
			wantErr: true,
		},
		{
			name:    "nil value",
			value:   reflect.ValueOf(nil),
			typ:     reflect.TypeOf(1),
			want:    reflect.ValueOf(nil),
			wantErr: false,
		},
		{
			name:    "nil value with nil type",
			value:   reflect.ValueOf(nil),
			typ:     nil,
			want:    reflect.ValueOf(nil),
			wantErr: false,
		},
		{
			name:    "nil value with non-nil type",
			value:   reflect.ValueOf(nil),
			typ:     reflect.TypeOf(1),
			want:    reflect.ValueOf(0),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &state{}
			got := s.validateType(tt.value, tt.typ)
			if (got.Interface() != tt.want.Interface()) || (got.IsValid() != tt.want.IsValid()) {
				t.Errorf("validateType() = %v, want %v", got, tt.want)
			}
		})
	}
}
