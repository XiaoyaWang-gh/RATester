package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrySetCustom_1(t *testing.T) {
	type args struct {
		val   string
		value reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				val:   "test",
				value: reflect.ValueOf(new(string)),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				val:   "test",
				value: reflect.ValueOf(new(int)),
			},
			want:    false,
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

			got, err := trySetCustom(tt.args.val, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("trySetCustom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("trySetCustom() = %v, want %v", got, tt.want)
			}
		})
	}
}
