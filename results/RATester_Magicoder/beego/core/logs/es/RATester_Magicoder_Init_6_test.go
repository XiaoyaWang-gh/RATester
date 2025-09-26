package es

import (
	"fmt"
	"testing"
)

func TestInit_6(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name    string
		el      *esLogger
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

			if err := tt.el.Init(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("esLogger.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
