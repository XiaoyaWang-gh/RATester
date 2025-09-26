package tcp

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestClientHelloInfo_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *clientHello
		wantErr bool
	}{
		{
			name:  "Test case 1",
			input: "test input 1",
			want: &clientHello{
				serverName: "test server name 1",
				protos:     []string{"test proto 1", "test proto 2"},
				isTLS:      true,
				peeked:     "test peeked 1",
			},
			wantErr: false,
		},
		{
			name:  "Test case 2",
			input: "test input 2",
			want: &clientHello{
				serverName: "test server name 2",
				protos:     []string{"test proto 3", "test proto 4"},
				isTLS:      false,
				peeked:     "test peeked 2",
			},
			wantErr: true,
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

			br := bufio.NewReader(strings.NewReader(tt.input))
			got, err := clientHelloInfo(br)
			if (err != nil) != tt.wantErr {
				t.Errorf("clientHelloInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clientHelloInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
