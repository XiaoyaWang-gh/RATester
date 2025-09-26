package metrics

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewProvider_1(t *testing.T) {
	type args struct {
		calculateHints bool
	}
	tests := []struct {
		name string
		args args
		want Provider
	}{
		{
			name: "Test case 1",
			args: args{
				calculateHints: true,
			},
			want: &Store{
				calculateHints: true,
				metrics:        make(map[string][]time.Duration),
				diffs:          make(map[string]*diff),
				cached:         make(map[string]int),
			},
		},
		{
			name: "Test case 2",
			args: args{
				calculateHints: false,
			},
			want: &Store{
				calculateHints: false,
				metrics:        make(map[string][]time.Duration),
				diffs:          make(map[string]*diff),
				cached:         make(map[string]int),
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

			if got := NewProvider(tt.args.calculateHints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
