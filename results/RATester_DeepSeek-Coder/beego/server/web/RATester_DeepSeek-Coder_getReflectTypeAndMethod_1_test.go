package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetReflectTypeAndMethod_1(t *testing.T) {
	type args struct {
		f interface{}
	}
	tests := []struct {
		name               string
		args               args
		wantControllerType reflect.Type
		wantMethod         string
		wantPanic          bool
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

			defer func() {
				if r := recover(); (r != nil) != tt.wantPanic {
					t.Errorf("getReflectTypeAndMethod() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			gotControllerType, gotMethod := getReflectTypeAndMethod(tt.args.f)
			if gotControllerType != tt.wantControllerType {
				t.Errorf("getReflectTypeAndMethod() gotControllerType = %v, want %v", gotControllerType, tt.wantControllerType)
			}
			if gotMethod != tt.wantMethod {
				t.Errorf("getReflectTypeAndMethod() gotMethod = %v, want %v", gotMethod, tt.wantMethod)
			}
		})
	}
}
