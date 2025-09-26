package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewRawPreparer_1(t *testing.T) {
	type args struct {
		rs *rawSet
	}
	tests := []struct {
		name    string
		args    args
		want    RawPreparer
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

			got, err := newRawPreparer(tt.args.rs)
			if (err != nil) != tt.wantErr {
				t.Errorf("newRawPreparer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRawPreparer() = %v, want %v", got, tt.want)
			}
		})
	}
}
