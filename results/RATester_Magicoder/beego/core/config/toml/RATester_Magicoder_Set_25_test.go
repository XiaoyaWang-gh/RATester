package toml

import (
	"fmt"
	"testing"
)

func TestSet_25(t *testing.T) {
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name    string
		c       *configContainer
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

			if err := tt.c.Set(tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("configContainer.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
