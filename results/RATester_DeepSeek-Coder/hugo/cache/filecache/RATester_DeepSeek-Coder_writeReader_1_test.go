package filecache

import (
	"fmt"
	"io"
	"testing"
)

func TestWriteReader_1(t *testing.T) {
	type args struct {
		id string
		r  io.Reader
	}
	tests := []struct {
		name    string
		c       *Cache
		args    args
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

			if err := tt.c.writeReader(tt.args.id, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Cache.writeReader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
