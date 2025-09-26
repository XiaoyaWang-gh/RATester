package nathole

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestDecodeMessageInto_1(t *testing.T) {
	type args struct {
		data []byte
		key  []byte
		m    msg.Message
	}
	tests := []struct {
		name    string
		args    args
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

			if err := DecodeMessageInto(tt.args.data, tt.args.key, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("DecodeMessageInto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
