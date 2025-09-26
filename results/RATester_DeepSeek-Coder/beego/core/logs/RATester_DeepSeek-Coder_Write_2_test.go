package logs

import (
	"fmt"
	"testing"
)

func TestWrite_2(t *testing.T) {
	tests := []struct {
		name    string
		bl      *BeeLogger
		input   []byte
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

			gotN, err := tt.bl.Write(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeeLogger.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("BeeLogger.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
