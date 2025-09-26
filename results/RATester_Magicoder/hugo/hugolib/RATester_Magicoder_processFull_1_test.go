package hugolib

import (
	"context"
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestprocessFull_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		l      logg.LevelLogger
		config *BuildCfg
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

			h := &HugoSites{}
			if err := h.processFull(tt.args.ctx, tt.args.l, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("HugoSites.processFull() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
