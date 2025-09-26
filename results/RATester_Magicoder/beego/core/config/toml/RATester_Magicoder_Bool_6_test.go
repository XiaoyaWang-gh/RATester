package toml

import (
	"fmt"
	"testing"
)

func TestBool_6(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		c       *configContainer
		args    args
		want    bool
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

			got, err := tt.c.Bool(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("configContainer.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("configContainer.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}
