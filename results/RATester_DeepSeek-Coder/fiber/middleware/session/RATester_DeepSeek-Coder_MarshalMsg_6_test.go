package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalMsg_6(t *testing.T) {
	testCases := []struct {
		name    string
		input   *data
		want    []byte
		wantErr bool
	}{
		{
			name: "Test case 1",
			input: &data{
				Data: map[string]any{
					"key1": "value1",
					"key2": 2,
					"key3": true,
				},
			},
			want: []byte{0x81, 0xa4, 0x44, 0x61, 0x74, 0x61, 0x83, 0xa4, 0x6b, 0x65, 0x79, 0x31, 0xa6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x31, 0xa4, 0x6b, 0x65, 0x79, 0x32, 0xc2, 0x02, 0xa4, 0x6b, 0x65, 0x79, 0x33, 0xc3, 0x01},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tc.input.MarshalMsg(nil)
			if (err != nil) != tc.wantErr {
				t.Errorf("MarshalMsg() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MarshalMsg() = %v, want %v", got, tc.want)
			}
		})
	}
}
