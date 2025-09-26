package toml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDIY_5(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		c       *configContainer
		args    args
		want    interface{}
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

			got, err := tt.c.DIY(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("configContainer.DIY() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configContainer.DIY() = %v, want %v", got, tt.want)
			}
		})
	}
}
