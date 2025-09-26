package session

import (
	"fmt"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestEncodeMsg_6(t *testing.T) {
	type args struct {
		en *msgp.Writer
	}
	tests := []struct {
		name    string
		z       *data
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			z: &data{
				Data: map[string]any{
					"key1": "value1",
					"key2": 123,
				},
			},
			args: args{
				en: &msgp.Writer{},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			z: &data{
				Data: map[string]any{
					"key1": "value1",
					"key2": make(chan int),
				},
			},
			args: args{
				en: &msgp.Writer{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			z := tt.z
			err := z.EncodeMsg(tt.args.en)
			if (err != nil) != tt.wantErr {
				t.Errorf("data.EncodeMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
