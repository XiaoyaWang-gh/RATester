package redis_sentinel

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInitOldStyle_4(t *testing.T) {
	type args struct {
		savePath string
	}
	tests := []struct {
		name string
		args args
		want Provider
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

			rp := &Provider{}
			rp.initOldStyle(tt.args.savePath)
			if !reflect.DeepEqual(*rp, tt.want) {
				t.Errorf("initOldStyle() = %v, want %v", *rp, tt.want)
			}
		})
	}
}
