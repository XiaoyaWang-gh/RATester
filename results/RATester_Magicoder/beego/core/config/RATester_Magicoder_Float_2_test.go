package config

import (
	"fmt"
	"testing"
)

func TestFloat_2(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
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
			got, err := c.Float(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Float() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Float() = %v, want %v", got, tt.want)
			}
		})
	}
}
