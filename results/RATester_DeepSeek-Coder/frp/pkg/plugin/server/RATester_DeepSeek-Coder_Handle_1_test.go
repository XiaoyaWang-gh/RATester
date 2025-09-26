package plugin

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestHandle_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		op      string
		content interface{}
	}
	tests := []struct {
		name    string
		p       *httpPlugin
		args    args
		want    *Response
		want1   interface{}
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

			got, got1, err := tt.p.Handle(tt.args.ctx, tt.args.op, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("httpPlugin.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpPlugin.Handle() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("httpPlugin.Handle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
