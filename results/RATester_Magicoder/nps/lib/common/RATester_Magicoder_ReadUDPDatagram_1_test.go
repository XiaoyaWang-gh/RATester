package common

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestReadUDPDatagram_1(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *UDPDatagram
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

			got, err := ReadUDPDatagram(tt.args.r)
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
