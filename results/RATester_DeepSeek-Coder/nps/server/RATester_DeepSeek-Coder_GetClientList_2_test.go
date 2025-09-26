package server

import (
	"fmt"
	"reflect"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestGetClientList_2(t *testing.T) {
	type args struct {
		start    int
		length   int
		search   string
		sort     string
		order    string
		clientId int
	}
	tests := []struct {
		name  string
		args  args
		want  []*file.Client
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

			got, got1 := GetClientList(tt.args.start, tt.args.length, tt.args.search, tt.args.sort, tt.args.order, tt.args.clientId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClientList() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetClientList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
