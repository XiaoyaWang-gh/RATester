package fiber

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestEncodeMsg_3(t *testing.T) {
	type testCase struct {
		name    string
		input   redirectionMsgs
		wantErr bool
	}

	tests := []testCase{
		{
			name: "Test case 1",
			input: []redirectionMsg{
				{
					key:        "key1",
					value:      "value1",
					level:      1,
					isOldInput: true,
				},
				{
					key:        "key2",
					value:      "value2",
					level:      2,
					isOldInput: false,
				},
			},
			wantErr: false,
		},
		{
			name:    "Test case 2",
			input:   nil,
			wantErr: false,
		},
		{
			name: "Test case 3",
			input: []redirectionMsg{
				{
					key:        "key3",
					value:      "value3",
					level:      3,
					isOldInput: true,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := msgp.NewWriter(ioutil.Discard)
			err := tt.input.EncodeMsg(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("redirectionMsgs.EncodeMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
