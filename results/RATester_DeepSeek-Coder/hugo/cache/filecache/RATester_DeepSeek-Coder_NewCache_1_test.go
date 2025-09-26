package filecache

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/BurntSushi/locker"
	"github.com/spf13/afero"
)

func TestNewCache_1(t *testing.T) {
	type args struct {
		fs              afero.Fs
		maxAge          time.Duration
		pruneAllRootDir string
	}
	tests := []struct {
		name string
		args args
		want *Cache
	}{
		{
			name: "TestNewCache",
			args: args{
				fs:              afero.NewMemMapFs(),
				maxAge:          1 * time.Hour,
				pruneAllRootDir: "/tmp",
			},
			want: &Cache{
				Fs:              afero.NewMemMapFs(),
				nlocker:         &lockTracker{Locker: locker.NewLocker(), seen: make(map[string]struct{})},
				maxAge:          1 * time.Hour,
				pruneAllRootDir: "/tmp",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewCache(tt.args.fs, tt.args.maxAge, tt.args.pruneAllRootDir)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
