package gin

import (
	"fmt"
	"testing"
)

func TestShouldBindUri_1(t *testing.T) {
	type testStruct struct {
		Param1 string `uri:"param1"`
		Param2 string `uri:"param2"`
	}

	tests := []struct {
		name    string
		context *Context
		wantErr bool
	}{
		{
			name: "should bind uri successfully",
			context: &Context{
				Params: Params{
					{
						Key:   "param1",
						Value: "value1",
					},
					{
						Key:   "param2",
						Value: "value2",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "should return error when params are not provided",
			context: &Context{
				Params: Params{},
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

			s := &testStruct{}
			if err := tt.context.ShouldBindUri(s); (err != nil) != tt.wantErr {
				t.Errorf("ShouldBindUri() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
