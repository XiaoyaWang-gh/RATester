package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLength_1(t *testing.T) {
	type args struct {
		item reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want    int
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

			got, err := length(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("length() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("length() = %v, want %v", got, tt.want)
			}
		})
	}
}
