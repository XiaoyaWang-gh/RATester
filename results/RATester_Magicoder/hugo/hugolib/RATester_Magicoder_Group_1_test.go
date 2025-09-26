package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroup_1(t *testing.T) {
	type args struct {
		key any
		in  any
	}
	tests := []struct {
		name    string
		args    args
		want    any
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

			p := &pageState{}
			got, err := p.Group(tt.args.key, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("pageState.Group() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pageState.Group() = %v, want %v", got, tt.want)
			}
		})
	}
}
