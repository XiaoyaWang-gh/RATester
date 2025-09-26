package hugolib

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestrenderShortcode_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		p       prerenderedShortcode
		args    args
		want    []byte
		want1   bool
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

			got, got1, err := tt.p.renderShortcode(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("renderShortcode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("renderShortcode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("renderShortcode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
