package opentelemetry

import (
	"fmt"
	"testing"
)

func TestSetup_1(t *testing.T) {
	type args struct {
		serviceName      string
		sampleRate       float64
		globalAttributes map[string]string
	}
	tests := []struct {
		name    string
		c       *Config
		args    args
		wantErr bool
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

			_, _, err := tt.c.Setup(tt.args.serviceName, tt.args.sampleRate, tt.args.globalAttributes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Setup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
