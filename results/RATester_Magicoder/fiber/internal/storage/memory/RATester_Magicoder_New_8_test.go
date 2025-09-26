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
			name: "TestNew_0",
			args: args{
				config: []Config{
					{
						GCInterval: 10 * time.Second,
					},
				},
			},
			want: &Storage{
				db:         make(map[string]entry),
				gcInterval: 10 * time.Second,
				done:       make(chan struct{}),
			},
		},
		{
			name: "TestNew_1",
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

			if got := New(tt.args.config...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
