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
			name: "Test with valid config",
			s:    &JLWriter{},
			args: args{
				config: `{"authorname": "test", "title": "test", "webhookurl": "http://test.com", "level": 1, "formatter": "testFormatter"}`,
			},
			wantErr: false,
		},
		{
			name: "Test with invalid config",
			s:    &JLWriter{},
			args: args{
				config: `{"authorname": "test", "title": "test", "webhookurl": "http://test.com", "level": 1, "formatter": "notExistFormatter"}`,
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

			s := &JLWriter{}
			if err := s.Init(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("JLWriter.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
