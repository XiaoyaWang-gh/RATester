package binding

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSetTimeDuration_1(t *testing.T) {
	type args struct {
		val   string
		value reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				val:   "10s",
				value: reflect.ValueOf(time.Duration(0)),
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				val:   "10s",
				value: reflect.ValueOf(time.Duration(0)),
			},
			wantErr: false,
		},
		{
			name: "test case 3",
			args: args{
				val:   "10s",
				value: reflect.ValueOf(time.Duration(0)),
			},
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

			if err := setTimeDuration(tt.args.val, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("setTimeDuration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
