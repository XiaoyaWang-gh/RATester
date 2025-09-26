package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetData_2(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		c    *Controller
		args args
		want map[interface{}]interface{}
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

			tt.c.SetData(tt.args.data)
			if !reflect.DeepEqual(tt.c.Data, tt.want) {
				t.Errorf("SetData() = %v, want %v", tt.c.Data, tt.want)
			}
		})
	}
}
