package task

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRunNextTasks_1(t *testing.T) {
	type fields struct {
		adminTaskList map[string]Tasker
		taskLock      sync.RWMutex
		stop          chan bool
		changed       chan bool
		started       bool
		wait          sync.WaitGroup
	}
	type args struct {
		sortList  *MapSorter
		effective time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			m := &taskManager{
				adminTaskList: tt.fields.adminTaskList,
				taskLock:      tt.fields.taskLock,
				stop:          tt.fields.stop,
				changed:       tt.fields.changed,
				started:       tt.fields.started,
				wait:          tt.fields.wait,
			}
			m.runNextTasks(tt.args.sortList, tt.args.effective)
		})
	}
}
