package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttryDecodeBodyInOrder_1(t *testing.T) {
	tests := []struct {
		name    string
		ctx     *DefaultCtx
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

			got, got1, err := tt.ctx.tryDecodeBodyInOrder(nil, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("tryDecodeBodyInOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tryDecodeBodyInOrder() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("tryDecodeBodyInOrder() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
