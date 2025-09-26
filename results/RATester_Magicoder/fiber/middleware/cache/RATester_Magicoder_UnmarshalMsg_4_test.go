package cache

import (
	"fmt"
	"testing"
)

func TestUnmarshalMsg_4(t *testing.T) {
	tests := []struct {
		name    string
		z       *item
		bts     []byte
		wantErr bool
	}{
		{
			name: "Test case 1",
			z: &item{
				headers:   map[string][]byte{"key": []byte("value")},
				body:      []byte("body"),
				ctype:     []byte("ctype"),
				cencoding: []byte("cencoding"),
				status:    200,
				exp:       1640995200,
				heapidx:   0,
			},
			bts:     []byte(`{"headers":{"key":["value"]},"body":["body"],"ctype":["ctype"],"cencoding":["cencoding"],"status":200,"exp":1640995200,"heapidx":0}`),
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

			_, err := tt.z.UnmarshalMsg(tt.bts)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
