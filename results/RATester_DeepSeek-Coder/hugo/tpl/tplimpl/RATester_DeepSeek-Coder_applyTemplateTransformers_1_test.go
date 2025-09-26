package tplimpl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestApplyTemplateTransformers_1(t *testing.T) {
	type args struct {
		t        *templateState
		lookupFn func(name string) *templateState
	}
	tests := []struct {
		name    string
		args    args
		want    *templateContext
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

			got, err := applyTemplateTransformers(tt.args.t, tt.args.lookupFn)
			if (err != nil) != tt.wantErr {
				t.Errorf("applyTemplateTransformers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyTemplateTransformers() = %v, want %v", got, tt.want)
			}
		})
	}
}
