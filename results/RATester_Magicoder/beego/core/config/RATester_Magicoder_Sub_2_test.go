package config

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestSub_2(t *testing.T) {
	type fields struct {
		reader func(ctx context.Context, key string) (string, error)
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Configer
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

			c := &BaseConfiger{
				reader: tt.fields.reader,
			}
			got, err := c.Sub(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseConfiger.Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseConfiger.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
