package hugolib

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestexpandShortcodeTokens_1(t *testing.T) {
	type args struct {
		ctx          context.Context
		source       []byte
		tokenHandler func(ctx context.Context, token string) ([]byte, error)
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
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

			got, err := expandShortcodeTokens(tt.args.ctx, tt.args.source, tt.args.tokenHandler)
			if (err != nil) != tt.wantErr {
				t.Errorf("expandShortcodeTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expandShortcodeTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}
