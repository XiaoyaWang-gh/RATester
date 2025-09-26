package pmux

import (
	"fmt"
	"testing"
)

func TestRead_4(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		pConn   *PortConn
		args    args
		wantN   int
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

			gotN, err := tt.pConn.Read(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("PortConn.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("PortConn.Read() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
