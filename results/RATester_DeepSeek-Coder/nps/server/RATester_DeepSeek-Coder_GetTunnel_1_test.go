package server

import (
	"fmt"
	"reflect"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestGetTunnel_1(t *testing.T) {
	type args struct {
		start    int
		length   int
		typeVal  string
		clientId int
		search   string
	}
	tests := []struct {
		name  string
		args  args
		want  []*file.Tunnel
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := GetTunnel(tt.args.start, tt.args.length, tt.args.typeVal, tt.args.clientId, tt.args.search)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTunnel() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetTunnel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
