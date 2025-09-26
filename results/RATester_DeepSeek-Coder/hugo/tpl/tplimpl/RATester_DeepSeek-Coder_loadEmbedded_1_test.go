package tplimpl

import (
	"fmt"
	"testing"
)

func TestLoadEmbedded_1(t *testing.T) {
	tests := []struct {
		name    string
		handler *templateHandler
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

			err := tt.handler.loadEmbedded()
			if (err != nil) != tt.wantErr {
				t.Errorf("templateHandler.loadEmbedded() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
