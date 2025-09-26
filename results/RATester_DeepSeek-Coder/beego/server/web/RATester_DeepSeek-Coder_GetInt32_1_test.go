package web

import (
	"fmt"
	"testing"
)

func TestGetInt32_1(t *testing.T) {
	type args struct {
		key string
		def []int32
	}
	tests := []struct {
		name    string
		args    args
		want    int32
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
			got, err := c.GetInt32(tt.args.key, tt.args.def...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}
