package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestStoreSyncMapToFile_1(t *testing.T) {
	type args struct {
		m        sync.Map
		filePath string
	}
	tests := []struct {
		name string
		args args
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

			storeSyncMapToFile(tt.args.m, tt.args.filePath)
		})
	}
}
