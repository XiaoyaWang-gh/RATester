package limit

import (
	"fmt"
	"testing"
)

func TestWrite_3(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		w       *Writer
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

			gotN, err := tt.w.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Writer.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Writer.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
