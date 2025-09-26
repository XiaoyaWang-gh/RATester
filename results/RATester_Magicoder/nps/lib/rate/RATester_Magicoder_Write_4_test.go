package rate

import (
	"fmt"
	"testing"
)

func TestWrite_4(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		s       *rateConn
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

			gotN, err := tt.s.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("rateConn.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("rateConn.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
