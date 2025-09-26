package resources

import (
	"fmt"
	"io"
	"testing"
)

func TestHashImage_1(t *testing.T) {
	type args struct {
		r io.ReadSeeker
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		want1   int64
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

			got, got1, err := hashImage(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("hashImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hashImage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("hashImage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
