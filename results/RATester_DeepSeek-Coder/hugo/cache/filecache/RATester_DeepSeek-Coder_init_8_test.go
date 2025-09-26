package filecache

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/spf13/afero"
)

func TestInit_8(t *testing.T) {
	type fields struct {
		Fs              afero.Fs
		maxAge          time.Duration
		pruneAllRootDir string
		nlocker         *lockTracker
		initOnce        sync.Once
		initErr         error
	}
	tests := []struct {
		name    string
		fields  fields
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

			c := &Cache{
				Fs:              tt.fields.Fs,
				maxAge:          tt.fields.maxAge,
				pruneAllRootDir: tt.fields.pruneAllRootDir,
				nlocker:         tt.fields.nlocker,
				initOnce:        tt.fields.initOnce,
				initErr:         tt.fields.initErr,
			}
			if err := c.init(); (err != nil) != tt.wantErr {
				t.Errorf("Cache.init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
