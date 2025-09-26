package cache

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestDecodeMsg_4(t *testing.T) {
	type test struct {
		name    string
		input   []byte
		want    item
		wantErr bool
	}

	tests := []test{
		{
			name:  "Test case 1",
			input: []byte(`{"headers":{"header1": "value1", "header2": "value2"}, "body": "body", "ctype": "ctype", "cencoding": "cencoding", "status": 200, "exp": 1640995200, "heapidx": 0}`),
			want: item{
				headers:   map[string][]byte{"header1": []byte("value1"), "header2": []byte("value2")},
				body:      []byte("body"),
				ctype:     []byte("ctype"),
				cencoding: []byte("cencoding"),
				status:    200,
				exp:       1640995200,
				heapidx:   0,
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			var got item
			err := got.DecodeMsg(msgp.NewReader(bytes.NewReader(tt.input)))

			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeMsg() got = %v, want %v", got, tt.want)
			}
		})
	}
}
