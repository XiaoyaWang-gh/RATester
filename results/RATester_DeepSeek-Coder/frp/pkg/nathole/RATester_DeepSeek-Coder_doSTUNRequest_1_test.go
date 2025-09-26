package nathole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDoSTUNRequest_1(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		c       *discoverConn
		args    args
		want    *stunResponse
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

			got, err := tt.c.doSTUNRequest(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("doSTUNRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doSTUNRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
