package logs

import (
	"fmt"
	"testing"
)

func TestInit_34(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name    string
		c       *consoleWriter
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			c:    &consoleWriter{},
			args: args{
				config: `{"formatter": "json", "level": 1, "color": true}`,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			c:    &consoleWriter{},
			args: args{
				config: `{"formatter": "invalid", "level": 1, "color": true}`,
			},
			wantErr: true,
		},
		{
			name: "Test case 3",
			c:    &consoleWriter{},
			args: args{
				config: `{"formatter": "", "level": 1, "color": true}`,
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			c:    &consoleWriter{},
			args: args{
				config: `{"formatter": "json", "level": 1, "color": true}`,
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

			if err := tt.c.Init(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("consoleWriter.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
