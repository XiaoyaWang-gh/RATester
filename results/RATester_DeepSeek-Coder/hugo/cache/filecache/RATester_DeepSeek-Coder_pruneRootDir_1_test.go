package filecache

import (
	"fmt"
	"testing"
)

func TestPruneRootDir_1(t *testing.T) {
	type args struct {
		force bool
	}
	tests := []struct {
		name    string
		c       *Cache
		args    args
		want    int
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

			got, err := tt.c.pruneRootDir(tt.args.force)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cache.pruneRootDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cache.pruneRootDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
