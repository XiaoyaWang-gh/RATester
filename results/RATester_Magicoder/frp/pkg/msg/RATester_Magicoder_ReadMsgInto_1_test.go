package msg

import (
	"fmt"
	"io"
	"testing"
)

func TestReadMsgInto_1(t *testing.T) {
	type args struct {
		c   io.Reader
		msg Message
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

			if err := ReadMsgInto(tt.args.c, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("ReadMsgInto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
