package nathole

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestEncodeMessage_1(t *testing.T) {
	type args struct {
		m   msg.Message
		key []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
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

			got, err := EncodeMessage(tt.args.m, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
