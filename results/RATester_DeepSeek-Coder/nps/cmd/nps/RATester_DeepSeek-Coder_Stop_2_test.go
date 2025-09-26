package main

import (
	"fmt"
	"testing"

	"github.com/kardianos/service"
)

func TestStop_2(t *testing.T) {
	type fields struct {
		exit chan struct{}
	}
	type args struct {
		s service.Service
	}
	tests := []struct {
		name    string
		fields  fields
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

			p := &nps{
				exit: tt.fields.exit,
			}
			if err := p.Stop(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("nps.Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
