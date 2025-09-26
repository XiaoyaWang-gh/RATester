package context

import (
	"fmt"
	"testing"
)

func TestYamlResp_1(t *testing.T) {
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
			name: "TestYamlResp_Success",
			ctx:  &Context{},
			args: args{
				data: map[string]interface{}{"key": "value"},
			},
			wantErr: false,
		},
		{
			name: "TestYamlResp_Fail",
			ctx:  &Context{},
			args: args{
				data: make(chan int),
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

			err := tt.ctx.YamlResp(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("YamlResp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
