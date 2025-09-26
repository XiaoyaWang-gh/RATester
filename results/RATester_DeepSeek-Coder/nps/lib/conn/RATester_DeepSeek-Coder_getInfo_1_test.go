package conn

import (
	"fmt"
	"testing"
)

func TestGetInfo_1(t *testing.T) {
	type args struct {
		t interface{}
	}
	tests := []struct {
		name    string
		s       *Conn
		args    args
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

			if err := tt.s.getInfo(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Conn.getInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
