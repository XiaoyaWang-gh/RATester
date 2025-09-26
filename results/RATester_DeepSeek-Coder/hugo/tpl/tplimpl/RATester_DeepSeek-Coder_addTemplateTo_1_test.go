package tplimpl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddTemplateTo_1(t *testing.T) {
	type args struct {
		info templateInfo
		to   *templateNamespace
	}
	tests := []struct {
		name    string
		args    args
		want    *templateState
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

			th := &templateHandler{}
			got, err := th.addTemplateTo(tt.args.info, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("templateHandler.addTemplateTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("templateHandler.addTemplateTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
