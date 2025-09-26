package yaml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsubMap_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		c       *ConfigContainer
		args    args
		want    map[string]interface{}
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

			got, err := tt.c.subMap(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigContainer.subMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigContainer.subMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
