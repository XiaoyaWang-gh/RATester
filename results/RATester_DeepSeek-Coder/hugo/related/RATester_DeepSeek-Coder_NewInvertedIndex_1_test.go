package related

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestNewInvertedIndex_1(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name string
		args args
		want *InvertedIndex
	}{
		{
			name: "Test with empty config",
			args: args{
				cfg: Config{},
			},
			want: &InvertedIndex{
				cfg:           Config{},
				index:         make(map[string]map[Keyword][]Document),
				indexDocCount: make(map[string]int),
			},
		},
		{
			name: "Test with non-empty config",
			args: args{
				cfg: Config{
					Threshold: 50,
					Indices: IndicesConfig{
						{Name: "index1"},
						{Name: "index2"},
					},
				},
			},
			want: &InvertedIndex{
				cfg: Config{
					Threshold: 50,
					Indices: IndicesConfig{
						{Name: "index1"},
						{Name: "index2"},
					},
				},
				index:         make(map[string]map[Keyword][]Document),
				indexDocCount: make(map[string]int),
				minWeight:     math.MaxInt64,
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

			got := NewInvertedIndex(tt.args.cfg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvertedIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
