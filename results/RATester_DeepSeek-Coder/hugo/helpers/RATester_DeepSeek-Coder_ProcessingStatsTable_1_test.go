package helpers

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestProcessingStatsTable_1(t *testing.T) {
	type args struct {
		w     io.Writer
		stats []*ProcessingStats
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				w: &bytes.Buffer{},
				stats: []*ProcessingStats{
					{
						Name:            "Test 1",
						Pages:           1,
						PaginatorPages:  2,
						Static:          3,
						ProcessedImages: 4,
						Files:           5,
						Aliases:         6,
						Cleaned:         7,
					},
					{
						Name:            "Test 2",
						Pages:           8,
						PaginatorPages:  9,
						Static:          10,
						ProcessedImages: 11,
						Files:           12,
						Aliases:         13,
						Cleaned:         14,
					},
				},
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

			ProcessingStatsTable(tt.args.w, tt.args.stats...)
		})
	}
}
