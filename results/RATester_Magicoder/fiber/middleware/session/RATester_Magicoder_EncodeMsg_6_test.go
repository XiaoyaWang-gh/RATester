package session

import (
	"fmt"
	"testing"
)

func TestEncodeMsg_6(t *testing.T) {
	tests := []struct {
		name    string
		z       *data
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
			wantErr: false,
		},
		{
			name: "Test case 2",
			z: &data{
				Data: map[string]any{
					"key1": "value1",
					"key2": 123,
					"key3": nil,
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			z: &data{
				Data: map[string]any{
					"key1": "value1",
					"key2": 123,
					"key3": make(chan int),
				},
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

			err := tt.z.EncodeMsg(nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
