package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalInputToField_1(t *testing.T) {

	type args struct {
		valueKind reflect.Kind
		val       string
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
				val:       "",
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

			got, err := unmarshalInputToField(tt.args.valueKind, tt.args.val, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalInputToField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("unmarshalInputToField() = %v, want %v", got, tt.want)
			}
		})
	}
}
