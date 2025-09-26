package config

import (
	"context"
	"fmt"
	"testing"
)

func TestString_8(t *testing.T) {
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
		want    string
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
			got, err := c.String(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseConfiger.String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BaseConfiger.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
