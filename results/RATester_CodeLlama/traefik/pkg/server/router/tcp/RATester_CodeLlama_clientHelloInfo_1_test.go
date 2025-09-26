package tcp

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestClientHelloInfo_1(t *testing.T) {
	type args struct {
		br *bufio.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *clientHello
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				br: bufio.NewReader(strings.NewReader("")),
			},
			want: &clientHello{
				peeked: "",
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

			got, err := clientHelloInfo(tt.args.br)
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
