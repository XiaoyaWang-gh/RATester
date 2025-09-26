package task

import (
	"fmt"
	"testing"
	"time"
)

func TestrunNextTasks_1(t *testing.T) {
	type args struct {
		sortList  *MapSorter
		effective time.Time
	}
	tests := []struct {
		name string
		args args
		want *taskManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &taskManager{}
			m.runNextTasks(tt.args.sortList, tt.args.effective)
		})
	}
}
