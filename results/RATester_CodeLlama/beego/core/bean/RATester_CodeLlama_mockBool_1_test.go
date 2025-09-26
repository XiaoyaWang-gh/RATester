package bean

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMockBool_1(t *testing.T) {
	type args struct {
		tagValue string
		pvv      reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				tagValue: "true",
				pvv:      reflect.ValueOf(true),
			},
			wantErr: false,
		},
		{
			name: "case 2",
			args: args{
				tagValue: "false",
				pvv:      reflect.ValueOf(false),
			},
			wantErr: false,
		},
		{
			name: "case 3",
			args: args{
				tagValue: "invalid",
				pvv:      reflect.ValueOf(false),
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

			if err := mockBool(tt.args.tagValue, tt.args.pvv); (err != nil) != tt.wantErr {
				t.Errorf("mockBool() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
