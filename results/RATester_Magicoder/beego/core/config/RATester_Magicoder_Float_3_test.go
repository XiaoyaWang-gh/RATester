package config

import (
	"fmt"
	"testing"
)

func TestFloat_3(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		c       *fakeConfigContainer
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

			got, err := tt.c.Float(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("fakeConfigContainer.Float() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fakeConfigContainer.Float() = %v, want %v", got, tt.want)
			}
		})
	}
}
