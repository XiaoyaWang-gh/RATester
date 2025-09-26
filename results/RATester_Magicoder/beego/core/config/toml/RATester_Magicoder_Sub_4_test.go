package toml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSub_4(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		c       *configContainer
		args    args
		want    *configContainer
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

			got, err := tt.c.Sub(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("configContainer.Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configContainer.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
