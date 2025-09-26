package conn

import (
	"fmt"
	"testing"
)

func TestReadLen_1(t *testing.T) {
	conn := &Conn{
		Conn: nil,
		Rb:   make([]byte, 10),
	}

	tests := []struct {
		name    string
		s       *Conn
		cLen    int
		buf     []byte
		wantN   int
		wantErr bool
	}{
		{
			name:    "Test case 1",
			s:       conn,
			cLen:    5,
			buf:     make([]byte, 10),
			wantN:   5,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			s:       conn,
			cLen:    15,
			buf:     make([]byte, 10),
			wantN:   0,
			wantErr: true,
		},
		{
			name:    "Test case 3",
			s:       conn,
			cLen:    -5,
			buf:     make([]byte, 10),
			wantN:   0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotN, err := tt.s.ReadLen(tt.cLen, tt.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ReadLen() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
