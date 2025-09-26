package logs

import (
	"fmt"
	"testing"
)

func TestInit_31(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name    string
		s       *JLWriter
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			s:    &JLWriter{},
			args: args{
				config: `{"authorname": "test", "title": "test", "webhookurl": "test", "redirecturl": "test", "imageurl": "test", "level": 1, "formatter": "test"}`,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			s:    &JLWriter{},
			args: args{
				config: `{"authorname": "test", "title": "test", "webhookurl": "test", "redirecturl": "test", "imageurl": "test", "level": 1, "formatter": "test"}`,
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

			if err := tt.s.Init(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("JLWriter.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
