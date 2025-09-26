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
		br      *bufio.Reader
		want    *clientHello
		wantErr bool
	}{
		{
			name: "Test case 1",
			br:   bufio.NewReader(strings.NewReader("test data")),
			want: &clientHello{
				peeked: "test data",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			br:   bufio.NewReader(strings.NewReader("")),
			want: &clientHello{
				peeked: "",
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			br:   bufio.NewReader(strings.NewReader("invalid data")),
			want: &clientHello{
				peeked: "invalid data",
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

			got, err := clientHelloInfo(tt.br)
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
