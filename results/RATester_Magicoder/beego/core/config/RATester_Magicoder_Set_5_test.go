package config

import (
	"fmt"
	"testing"
)

func TestSet_5(t *testing.T) {
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name    string
		c       *IniConfigContainer
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
				t.Errorf("IniConfigContainer.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
