package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestStoreTasksToJsonFile_1(t *testing.T) {
	type fields struct {
		Tasks            sync.Map
		Hosts            sync.Map
		HostsTmp         sync.Map
		Clients          sync.Map
		RunPath          string
		ClientIncreaseId int32
		TaskIncreaseId   int32
		HostIncreaseId   int32
		TaskFilePath     string
		HostFilePath     string
		ClientFilePath   string
	}
	tests := []struct {
		name   string
		fields fields
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

			s := &JsonDb{
				Tasks:            tt.fields.Tasks,
				Hosts:            tt.fields.Hosts,
				HostsTmp:         tt.fields.HostsTmp,
				Clients:          tt.fields.Clients,
				RunPath:          tt.fields.RunPath,
				ClientIncreaseId: tt.fields.ClientIncreaseId,
				TaskIncreaseId:   tt.fields.TaskIncreaseId,
				HostIncreaseId:   tt.fields.HostIncreaseId,
				TaskFilePath:     tt.fields.TaskFilePath,
				HostFilePath:     tt.fields.HostFilePath,
				ClientFilePath:   tt.fields.ClientFilePath,
			}
			s.StoreTasksToJsonFile()
		})
	}
}
