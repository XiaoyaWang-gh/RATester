package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewTask_1(t *testing.T) {
	db := &DbUtils{
		JsonDb: &JsonDb{
			Tasks: sync.Map{},
		},
	}

	testCases := []struct {
		name    string
		task    *Tunnel
		wantErr bool
	}{
		{
			name: "success",
			task: &Tunnel{
				Id:       1,
				Port:     8080,
				Mode:     "secret",
				Password: "password",
			},
			wantErr: false,
		},
		{
			name: "error",
			task: &Tunnel{
				Id:       2,
				Port:     8081,
				Mode:     "secret",
				Password: "password",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := db.NewTask(tc.task)
			if (err != nil) != tc.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
