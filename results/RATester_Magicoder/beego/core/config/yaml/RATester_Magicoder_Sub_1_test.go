package yaml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSub_1(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ConfigContainer
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

			c := &ConfigContainer{
				data: tt.fields.data,
			}
			got, err := c.Sub(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigContainer.Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigContainer.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
