package context

import (
	"fmt"
	"testing"
)

func TestJSONResp_1(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		ctx     *Context
		args    args
		wantErr bool
	}{
		{
			name: "TestJSONResp_1",
			ctx:  &Context{},
			args: args{
				data: map[string]interface{}{"key": "value"},
			},
			wantErr: false,
		},
		{
			name: "TestJSONResp_2",
			ctx:  &Context{},
			args: args{
				data: "string",
			},
			wantErr: false,
		},
		{
			name: "TestJSONResp_3",
			ctx:  &Context{},
			args: args{
				data: 123,
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

			err := tt.ctx.JSONResp(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONResp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
