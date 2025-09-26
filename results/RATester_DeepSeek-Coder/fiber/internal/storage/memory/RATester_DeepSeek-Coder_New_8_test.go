package memory

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNew_8(t *testing.T) {
	type args struct {
		config []Config
	}
	tests := []struct {
		name string
		args args
		want *Storage
	}{
		{
			name: "Test with default config",
			args: args{
				config: []Config{},
			},
			want: &Storage{
				db:         make(map[string]entry),
				gcInterval: 10 * time.Second,
				done:       make(chan struct{}),
			},
		},
		{
			name: "Test with custom config",
			args: args{
				config: []Config{
					{
						GCInterval: 20 * time.Second,
					},
				},
			},
			want: &Storage{
				db:         make(map[string]entry),
				gcInterval: 20 * time.Second,
				done:       make(chan struct{}),
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

			got := New(tt.args.config...)
			if !reflect.DeepEqual(got.db, tt.want.db) {
				t.Errorf("New() = %v, want %v", got.db, tt.want.db)
			}
			if got.gcInterval != tt.want.gcInterval {
				t.Errorf("New() = %v, want %v", got.gcInterval, tt.want.gcInterval)
			}
			if reflect.TypeOf(got.done) != reflect.TypeOf(tt.want.done) {
				t.Errorf("New() = %T, want %T", got.done, tt.want.done)
			}
		})
	}
}
