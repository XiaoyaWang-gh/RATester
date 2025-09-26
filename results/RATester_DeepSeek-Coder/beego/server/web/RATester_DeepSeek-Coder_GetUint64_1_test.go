package web

import (
	"fmt"
	"testing"
)

func TestGetUint64_1(t *testing.T) {
	type args struct {
		key string
		def []uint64
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
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
			got, err := c.GetUint64(tt.args.key, tt.args.def...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
