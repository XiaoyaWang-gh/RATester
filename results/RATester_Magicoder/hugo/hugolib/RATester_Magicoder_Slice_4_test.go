package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlice_4(t *testing.T) {
	type args struct {
		items any
	}
	tests := []struct {
		name    string
		p       *pageState
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

			got, err := tt.p.Slice(tt.args.items)
			if (err != nil) != tt.wantErr {
				t.Errorf("pageState.Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pageState.Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
