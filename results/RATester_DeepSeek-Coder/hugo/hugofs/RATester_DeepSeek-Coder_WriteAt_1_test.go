package hugofs

import (
	"fmt"
	"testing"
)

func TestWriteAt_1(t *testing.T) {
	type args struct {
		p   []byte
		off int64
	}
	tests := []struct {
		name    string
		f       *noOpRegularFileOps
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

			f := &noOpRegularFileOps{}
			gotN, err := f.WriteAt(tt.args.p, tt.args.off)
			if (err != nil) != tt.wantErr {
				t.Errorf("noOpRegularFileOps.WriteAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("noOpRegularFileOps.WriteAt() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
