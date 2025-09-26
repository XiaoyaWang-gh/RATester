package dynacache

import (
	"fmt"
	"testing"
)

func TestCalculateMaxSizePerPartition_1(t *testing.T) {
	type args struct {
		maxItemsTotal       int
		totalWeightQuantity int
		numPartitions       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				maxItemsTotal:       100,
				totalWeightQuantity: 100,
				numPartitions:       10,
			},
			want: 10,
		},
		{
			name: "test case 2",
			args: args{
				maxItemsTotal:       100,
				totalWeightQuantity: 100,
				numPartitions:       100,
			},
			want: 1,
		},
		{
			name: "test case 3",
			args: args{
				maxItemsTotal:       100,
				totalWeightQuantity: 100,
				numPartitions:       1000,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := calculateMaxSizePerPartition(tt.args.maxItemsTotal, tt.args.totalWeightQuantity, tt.args.numPartitions); got != tt.want {
				t.Errorf("calculateMaxSizePerPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
