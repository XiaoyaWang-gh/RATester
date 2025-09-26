package commands

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetErrorWithContext_1(t *testing.T) {
	type fields struct {
		r *rootCommand
		c *serverCommand
	}
	tests := []struct {
		name   string
		fields fields
		want   any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &serverCommand{
				r:            tt.fields.r,
				commands:     tt.fields.c.commands,
				hugoBuilder:  tt.fields.c.hugoBuilder,
				quit:         tt.fields.c.quit,
				serverPorts:  tt.fields.c.serverPorts,
				doLiveReload: tt.fields.c.doLiveReload,
				// other fields...
			}
			if got := c.getErrorWithContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serverCommand.getErrorWithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
