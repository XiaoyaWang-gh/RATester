package proxy

import (
	"fmt"
	"testing"
)

func TestAdd_4(t *testing.T) {
	type args struct {
		name string
		pxy  Proxy
	}
	tests := []struct {
		name    string
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

			pm := &Manager{
				pxys: make(map[string]Proxy),
			}
			if err := pm.Add(tt.args.name, tt.args.pxy); (err != nil) != tt.wantErr {
				t.Errorf("Manager.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
