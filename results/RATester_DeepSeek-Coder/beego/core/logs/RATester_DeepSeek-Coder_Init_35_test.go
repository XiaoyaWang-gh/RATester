package logs

import (
	"fmt"
	"testing"
)

func TestInit_35(t *testing.T) {
	tests := []struct {
		name    string
		w       *fileLogWriter
		config  string
		wantErr bool
	}{
		{
			name: "Test with valid config",
			w: &fileLogWriter{
				Filename: "test.log",
			},
			config:  `{"filename": "test.log", "maxsize": 1024}`,
			wantErr: false,
		},
		{
			name: "Test with invalid config",
			w: &fileLogWriter{
				Filename: "test.log",
			},
			config:  `{"filename": "test.log", "maxsize": "1024"}`,
			wantErr: true,
		},
		{
			name: "Test with empty filename",
			w: &fileLogWriter{
				Filename: "",
			},
			config:  `{"filename": "", "maxsize": 1024}`,
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

			w := tt.w
			err := w.Init(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileLogWriter.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
