package hydratation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetMap_1(t *testing.T) {
	type args struct {
		field reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				field: reflect.ValueOf([]string{"test"}),
			},
			wantErr: true,
		},
		{
			name: "test case 2",
			args: args{
				field: reflect.ValueOf([]string{"test"}),
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

			if err := setMap(tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("setMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
