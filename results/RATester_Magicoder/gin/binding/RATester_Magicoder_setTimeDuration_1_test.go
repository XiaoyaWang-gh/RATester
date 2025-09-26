package binding

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestsetTimeDuration_1(t *testing.T) {
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
			name: "Test case 1",
			args: args{
				val:   "1h",
				value: reflect.ValueOf(time.Duration(0)),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				val:   "invalid",
				value: reflect.ValueOf(time.Duration(0)),
			},
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

			if err := setTimeDuration(tt.args.val, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("setTimeDuration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
