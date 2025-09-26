package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNormalizeMounts_1(t *testing.T) {
	type args struct {
		owner  *moduleAdapter
		mounts []Mount
	}
	tests := []struct {
		name    string
		args    args
		want    []Mount
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

			c := &collector{
				Client: &Client{},
			}
			got, err := c.normalizeMounts(tt.args.owner, tt.args.mounts)
			if (err != nil) != tt.wantErr {
				t.Errorf("normalizeMounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalizeMounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
