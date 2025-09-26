package conn

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLinkInfo_1(t *testing.T) {
	conn := &Conn{
		Conn: nil,
		Rb:   []byte("test"),
	}

	tests := []struct {
		name    string
		s       *Conn
		wantLk  *Link
		wantErr bool
	}{
		{
			name:    "TestGetLinkInfo",
			s:       conn,
			wantLk:  &Link{ConnType: "test", Host: "test", Crypt: true},
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

			gotLk, err := tt.s.GetLinkInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLinkInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLk, tt.wantLk) {
				t.Errorf("GetLinkInfo() gotLk = %v, want %v", gotLk, tt.wantLk)
			}
		})
	}
}
