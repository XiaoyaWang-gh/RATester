package ledis

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_3(t *testing.T) {
	type fields struct {
		maxlifetime int64
		SavePath    string `json:"save_path"`
		Db          int    `json:"db"`
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
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

			p := &Provider{
				maxlifetime: tt.fields.maxlifetime,
				SavePath:    tt.fields.SavePath,
				Db:          tt.fields.Db,
			}
			if got := p.SessionAll(tt.args.ctx); got != tt.want {
				t.Errorf("SessionAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
