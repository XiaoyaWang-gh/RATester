package msg

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestReadMsg_1(t *testing.T) {
	type args struct {
		c io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantMsg Message
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

			gotMsg, err := ReadMsg(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("ReadMsg() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}
