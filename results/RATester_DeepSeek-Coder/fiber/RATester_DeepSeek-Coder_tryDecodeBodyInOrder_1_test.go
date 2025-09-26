package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTryDecodeBodyInOrder_1(t *testing.T) {
	type args struct {
		originalBody *[]byte
		encodings    []string
	}
	tests := []struct {
		name    string
		c       *DefaultCtx
		args    args
		want    []byte
		want1   uint8
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

			got, got1, err := tt.c.tryDecodeBodyInOrder(tt.args.originalBody, tt.args.encodings)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultCtx.tryDecodeBodyInOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultCtx.tryDecodeBodyInOrder() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DefaultCtx.tryDecodeBodyInOrder() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
