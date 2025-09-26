package commands

import (
	"fmt"
	"net"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateEndpoint_1(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		f       *fileServer
		args    args
		want    *http.ServeMux
		want1   net.Listener
		want2   string
		want3   string
		wantErr bool
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

			got, got1, got2, got3, err := tt.f.createEndpoint(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileServer.createEndpoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileServer.createEndpoint() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("fileServer.createEndpoint() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("fileServer.createEndpoint() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("fileServer.createEndpoint() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
