package session

import (
	"fmt"
	"strings"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestDecodeMsg_6(t *testing.T) {
	type testCase struct {
		name    string
		input   string
		wantErr bool
	}

	tests := []testCase{
		{
			name: "valid data",
			input: `{
				"Data": {
					"key1": "value1",
					"key2": 2,
					"key3": 3.0,
					"key4": true,
					"key5": null
				}
			}`,
			wantErr: false,
		},
		{
			name: "invalid data",
			input: `{
				"Data": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
					"key5": "value5"
				}
			}`,
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

			r := msgp.NewReader(strings.NewReader(tt.input))
			d := &data{}
			err := d.DecodeMsg(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
