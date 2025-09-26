package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFromStringAndExt_1(t *testing.T) {
	type args struct {
		t   string
		ext []string
	}
	tests := []struct {
		name    string
		args    args
		want    Type
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

			got, err := FromStringAndExt(tt.args.t, tt.args.ext...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromStringAndExt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromStringAndExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
