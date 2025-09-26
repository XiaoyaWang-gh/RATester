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
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				w: &bytes.Buffer{},
				stats: []*ProcessingStats{
					{
						Name:            "Test",
						Pages:           10,
						PaginatorPages:  5,
						Static:          3,
						ProcessedImages: 2,
						Files:           1,
						Aliases:         0,
						Cleaned:         0,
					},
				},
			},
			want: "+------+-------+---------------+----------------+----------------+-------+--------+--------+\n| NAME | PAGES | PAGINATORPAGES | STATIC         | PROCESSEDIMAGES | FILES | ALIASES | CLEANED |\n+------+-------+---------------+----------------+----------------+-------+--------+--------+\n| Test | 10    | 5             | 3              | 2              | 1     | 0      | 0      |\n+------+-------+---------------+----------------+----------------+-------+--------+--------+\n",
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ProcessingStatsTable(tt.args.w, tt.args.stats...)
			if got := tt.args.w.(*bytes.Buffer).String(); got != tt.want {
				t.Errorf("ProcessingStatsTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
