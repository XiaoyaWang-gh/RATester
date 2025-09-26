package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeGob_1(t *testing.T) {
	type args struct {
		obj map[interface{}]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				obj: map[interface{}]interface{}{
					"key1": "value1",
					"key2": 123,
				},
			},
			want:    []byte{0x1, 0x2, 0x3, 0x4}, // Replace with the actual encoded bytes
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				obj: map[interface{}]interface{}{
					"key1": "value1",
					"key2": make(chan int), // Invalid type
				},
			},
			want:    nil,
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

			got, err := EncodeGob(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeGob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeGob() = %v, want %v", got, tt.want)
			}
		})
	}
}
