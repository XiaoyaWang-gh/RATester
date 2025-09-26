package dynacache

import (
	"fmt"
	"testing"
)

func TestcalculateMaxSizePerPartition_1(t *testing.T) {
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
			name: "Test case 1",
			args: args{
				maxItemsTotal:       100,
				totalWeightQuantity: 50,
				numPartitions:       2,
			},
			want: 50,
		},
		{
			name: "Test case 2",
			args: args{
				maxItemsTotal:       100,
				totalWeightQuantity: 100,
				numPartitions:       1,
			},
			want: 100,
		},
		{
			name: "Test case 3",
			args: args{
				maxItemsTotal:       100,
				totalWeightQuantity: 0,
				numPartitions:       1,
			},
			want: 0,
		},
		{
			name: "Test case 4",
			args: args{
				maxItemsTotal:       100,
				totalWeightQuantity: 100,
				numPartitions:       0,
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
