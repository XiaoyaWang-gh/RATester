package web

import (
	"fmt"
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
		})
	}
}
