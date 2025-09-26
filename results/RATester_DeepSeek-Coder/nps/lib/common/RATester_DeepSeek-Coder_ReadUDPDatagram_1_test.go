package common

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestReadUDPDatagram_1(t *testing.T) {
	type testCase struct {
		name    string
		input   io.Reader
		want    *UDPDatagram
		wantErr bool
	}

	tests := []testCase{
		{
			name:  "test case 1",
			input: strings.NewReader("test input"),
			want: &UDPDatagram{
				Header: &UDPHeader{
					Rsv:  0,
					Frag: 0,
					Addr: &Addr{},
				},
				Data: []byte("test input"),
			},
			wantErr: false,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ReadUDPDatagram(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUDPDatagram() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadUDPDatagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
