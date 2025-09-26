package config

import (
	"fmt"
	"testing"
)

func TestBool_4(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
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

			c := &BaseConfiger{}
			got, err := c.Bool(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseConfiger.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BaseConfiger.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}
