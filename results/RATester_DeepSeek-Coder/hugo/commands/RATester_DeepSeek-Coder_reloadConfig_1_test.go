package commands

import (
	"fmt"
	"sync"
	"testing"
)

func TestReloadConfig_1(t *testing.T) {
	type fields struct {
		r              *rootCommand
		confmu         sync.Mutex
		conf           *commonConfig
		onConfigLoaded func(reloaded bool) error
	}
	tests := []struct {
		name    string
		fields  fields
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

			c := &hugoBuilder{
				r:              tt.fields.r,
				confmu:         tt.fields.confmu,
				conf:           tt.fields.conf,
				onConfigLoaded: tt.fields.onConfigLoaded,
			}
			if err := c.reloadConfig(); (err != nil) != tt.wantErr {
				t.Errorf("hugoBuilder.reloadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
