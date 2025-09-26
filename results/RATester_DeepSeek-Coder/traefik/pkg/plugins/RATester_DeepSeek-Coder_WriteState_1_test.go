package plugins

import (
	"fmt"
	"testing"
)

func TestWriteState_1(t *testing.T) {
	type args struct {
		plugins map[string]Descriptor
	}

	tests := []struct {
		name    string
		c       *Client
		args    args
		wantErr bool
	}{
		{
			name: "TestWriteState_Success",
			c: &Client{
				stateFile: "testdata/state.json",
			},
			args: args{
				plugins: map[string]Descriptor{
					"plugin1": {
						ModuleName: "module1",
						Version:    "v1.0.0",
					},
					"plugin2": {
						ModuleName: "module2",
						Version:    "v2.0.0",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "TestWriteState_Fail",
			c: &Client{
				stateFile: "/non-existing-dir/state.json",
			},
			args: args{
				plugins: map[string]Descriptor{
					"plugin1": {
						ModuleName: "module1",
						Version:    "v1.0.0",
					},
					"plugin2": {
						ModuleName: "module2",
						Version:    "v2.0.0",
					},
				},
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

			c := tt.c
			err := c.WriteState(tt.args.plugins)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.WriteState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
