package hugio

import (
	"fmt"
	"testing"
)

func TestWrite_1(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		h       *HasBytesWriter
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

			gotN, err := tt.h.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("HasBytesWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("HasBytesWriter.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
