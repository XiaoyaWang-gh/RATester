package limit

import (
	"fmt"
	"testing"
)

func TestRead_2(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		r       *Reader
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

			gotN, err := tt.r.Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Reader.Read() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
