package logger

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateTagMap_1(t *testing.T) {
	type args struct {
		cfg *Config
	}
	tests := []struct {
		name string
		args args
		want map[string]LogFunc
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

			if got := createTagMap(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTagMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
