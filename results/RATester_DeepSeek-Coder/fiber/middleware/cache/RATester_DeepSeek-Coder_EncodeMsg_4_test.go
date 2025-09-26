package cache

import (
	"fmt"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestEncodeMsg_4(t *testing.T) {
	type args struct {
		en *msgp.Writer
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

			err := tt.z.EncodeMsg(tt.args.en)
			if (err != nil) != tt.wantErr {
				t.Errorf("item.EncodeMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
