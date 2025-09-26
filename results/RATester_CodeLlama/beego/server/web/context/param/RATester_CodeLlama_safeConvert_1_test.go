package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSafeConvert_1(t *testing.T) {
	t.Parallel()
	type args struct {
		value reflect.Value
		t     reflect.Type
	}
	tests := []struct {
		name    string
		args    args
		want    reflect.Value
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				value: reflect.ValueOf(1),
				t:     reflect.TypeOf(1),
			},
			want:    reflect.ValueOf(1),
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				value: reflect.ValueOf(1),
				t:     reflect.TypeOf(""),
			},
			want:    reflect.ValueOf(""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := safeConvert(tt.args.value, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("safeConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("safeConvert() got = %v, want %v", got, tt.want)
			}
		})
	}
}
