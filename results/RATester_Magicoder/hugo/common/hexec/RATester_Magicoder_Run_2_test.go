package hexec

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestRun_2(t *testing.T) {
	type fields struct {
		name   string
		c      *exec.Cmd
		outerr *bytes.Buffer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				name: "test",
				c: &exec.Cmd{
					Path: "test",
					Args: []string{"test"},
				},
				outerr: bytes.NewBufferString(""),
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			fields: fields{
				name: "test",
				c: &exec.Cmd{
					Path: "test",
					Args: []string{"test"},
				},
				outerr: bytes.NewBufferString("not found"),
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

			c := &cmdWrapper{
				name:   tt.fields.name,
				c:      tt.fields.c,
				outerr: tt.fields.outerr,
			}
			if err := c.Run(); (err != nil) != tt.wantErr {
				t.Errorf("cmdWrapper.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
