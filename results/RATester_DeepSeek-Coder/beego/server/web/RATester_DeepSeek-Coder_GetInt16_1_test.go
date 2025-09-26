package web

import (
	"fmt"
	"testing"
)

func TestGetInt16_1(t *testing.T) {
	type args struct {
		key string
		def []int16
	}
	tests := []struct {
		name    string
		args    args
		want    int16
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

			c := &Controller{}
			got, err := c.GetInt16(tt.args.key, tt.args.def...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}
