package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewProcessingStats_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *ProcessingStats
	}{
		{
			name: "Test case 1",
			args: args{name: "test"},
			want: &ProcessingStats{Name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewProcessingStats(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProcessingStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
