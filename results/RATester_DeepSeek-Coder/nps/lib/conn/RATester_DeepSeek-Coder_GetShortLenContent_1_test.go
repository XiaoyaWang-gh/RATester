package conn

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestGetShortLenContent_1(t *testing.T) {
	type fields struct {
		Conn net.Conn
		Rb   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantB   []byte
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

			s := &Conn{
				Conn: tt.fields.Conn,
				Rb:   tt.fields.Rb,
			}
			gotB, err := s.GetShortLenContent()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShortLenContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("GetShortLenContent() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
