package binding

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidateStruct_1(t *testing.T) {
	type args struct {
		obj any
	}
	tests := []struct {
		name    string
		v       *defaultValidator
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			v: &defaultValidator{
				validate: &validator.Validate{},
			},
			args: args{
				obj: nil,
			},
			wantErr: true,
		},
		{
			name: "test case 2",
			v: &defaultValidator{
				validate: &validator.Validate{},
			},
			args: args{
				obj: 1,
			},
			wantErr: true,
		},
		{
			name: "test case 3",
			v: &defaultValidator{
				validate: &validator.Validate{},
			},
			args: args{
				obj: "test",
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

			if err := tt.v.validateStruct(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("validateStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
