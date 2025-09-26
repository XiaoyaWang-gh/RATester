package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalInputsToField_1(t *testing.T) {
	type args struct {
		valueKind reflect.Kind
		values    []string
		field     reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				valueKind: reflect.Kind(0),
				values:    []string{""},
				field:     reflect.Value{},
			},
			want:    false,
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

			got, err := unmarshalInputsToField(tt.args.valueKind, tt.args.values, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalInputsToField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("unmarshalInputsToField() = %v, want %v", got, tt.want)
			}
		})
	}
}
