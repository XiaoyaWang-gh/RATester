package cache

import (
	"fmt"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestDecodeMsg_4(t *testing.T) {
	type args struct {
		dc *msgp.Reader
	}
	tests := []struct {
		name    string
		z       *item
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

			if err := tt.z.DecodeMsg(tt.args.dc); (err != nil) != tt.wantErr {
				t.Errorf("item.DecodeMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
