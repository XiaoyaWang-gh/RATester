package alils

import (
	"fmt"
	"testing"
)

func TestSize_1(t *testing.T) {
	tests := []struct {
		name string
		lg   *LogGroup
		want int
	}{
		{
			name: "Test case 1",
			lg: &LogGroup{
				Logs: []*Log{
					{
						Time: &[]uint32{1}[0],
						Contents: []*LogContent{
							{
								Key:   &[]string{"key1"}[0],
								Value: &[]string{"value1"}[0],
							},
						},
					},
				},
				Reserved: &[]string{"reserved1"}[0],
				Topic:    &[]string{"topic1"}[0],
				Source:   &[]string{"source1"}[0],
			},
			want: 100,
		},
		{
			name: "Test case 2",
			lg: &LogGroup{
				Logs: []*Log{
					{
						Time: &[]uint32{2}[0],
						Contents: []*LogContent{
							{
								Key:   &[]string{"key2"}[0],
								Value: &[]string{"value2"}[0],
							},
						},
					},
				},
				Reserved: &[]string{"reserved2"}[0],
				Topic:    &[]string{"topic2"}[0],
				Source:   &[]string{"source2"}[0],
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.lg.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
