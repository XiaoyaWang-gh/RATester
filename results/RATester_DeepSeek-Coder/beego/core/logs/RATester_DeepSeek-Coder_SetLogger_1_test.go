package logs

import (
	"fmt"
	"testing"
)

func TestSetLogger_1(t *testing.T) {
	tests := []struct {
		name        string
		adapterName string
		configs     []string
		wantErr     bool
	}{
		{
			name:        "Test case 1",
			adapterName: "file",
			configs:     []string{"test.log"},
			wantErr:     false,
		},
		{
			name:        "Test case 2",
			adapterName: "unknown",
			configs:     []string{"test.log"},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			bl := &BeeLogger{}
			err := bl.SetLogger(tt.adapterName, tt.configs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeeLogger.SetLogger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
