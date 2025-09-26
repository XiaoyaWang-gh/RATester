package main

import (
	"fmt"
	"go/types"
	"testing"
)

func TestWriteStruct_1(t *testing.T) {
	type args struct {
		name    string
		obj     *types.Struct
		rootPkg string
		elt     *File
	}

	tests := []struct {
		name string
		c    Centrifuge
		args args
		want string
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

			if got := tt.c.writeStruct(tt.args.name, tt.args.obj, tt.args.rootPkg, tt.args.elt); got != tt.want {
				t.Errorf("Centrifuge.writeStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
