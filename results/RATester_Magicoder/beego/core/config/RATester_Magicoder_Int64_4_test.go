package config

import (
	"fmt"
	"testing"
)

func TestInt64_4(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		c       *IniConfigContainer
		args    args
		want    int64
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

			got, err := tt.c.Int64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("IniConfigContainer.Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IniConfigContainer.Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}
