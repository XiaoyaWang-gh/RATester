package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDIY_4(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		c       *fakeConfigContainer
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
				t.Errorf("fakeConfigContainer.DIY() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fakeConfigContainer.DIY() = %v, want %v", got, tt.want)
			}
		})
	}
}
