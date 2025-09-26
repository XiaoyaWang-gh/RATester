package logs

import (
	"fmt"
	"testing"
)

func TestInit_33(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name    string
		s       *SLACKWriter
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid config",
			s:    &SLACKWriter{},
			args: args{
				config: `{"webhookurl": "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX", "level": 7, "formatter": "text"}`,
			},
			wantErr: false,
		},
		{
			name: "Test with invalid config",
			s:    &SLACKWriter{},
			args: args{
				config: `{"webhookurl": "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX", "level": 7, "formatter": "invalid"}`,
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

			s := &SLACKWriter{}
			if err := s.Init(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("SLACKWriter.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
