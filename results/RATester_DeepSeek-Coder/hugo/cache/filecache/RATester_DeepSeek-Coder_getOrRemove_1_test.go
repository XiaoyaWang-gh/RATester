package filecache

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/gohugoio/hugo/common/hugio"
	"github.com/spf13/afero"
)

func TestGetOrRemove_1(t *testing.T) {
	type fields struct {
		Fs              afero.Fs
		maxAge          time.Duration
		pruneAllRootDir string
		nlocker         *lockTracker
		initOnce        sync.Once
		initErr         error
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   hugio.ReadSeekCloser
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
			if got := c.getOrRemove(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOrRemove() = %v, want %v", got, tt.want)
			}
		})
	}
}
