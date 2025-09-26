package alils

import (
	"fmt"
	"testing"
)

func TestSize_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lg := &LogGroupList{
		LogGroups: []*LogGroup{
			{
				Logs: []*Log{
					{
						Time: &[]uint32{1}[0],
						Contents: []*LogContent{
							{
								Key:   &[]string{"key"}[0],
								Value: &[]string{"value"}[0],
							},
						},
					},
				},
			},
		},
	}
	expected := 100
	actual := lg.Size()
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
