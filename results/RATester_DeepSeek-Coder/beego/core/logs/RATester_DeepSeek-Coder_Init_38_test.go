package logs

import (
	"fmt"
	"testing"
)

func TestInit_38(t *testing.T) {
	tests := []struct {
		name    string
		s       *SMTPWriter
		config  string
		wantErr bool
	}{
		{
			name: "Test with valid config",
			s: &SMTPWriter{
				Username: "test",
				Password: "test",
				Host:     "test",
			},
			config:  `{"username":"test","password":"test","host":"test"}`,
			wantErr: false,
		},
		{
			name: "Test with invalid config",
			s: &SMTPWriter{
				Username: "test",
				Password: "test",
				Host:     "test",
			},
			config:  `{"username":"test","password":"test","host":"test","formatter":"json"}`,
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

			if err := tt.s.Init(tt.config); (err != nil) != tt.wantErr {
				t.Errorf("SMTPWriter.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
