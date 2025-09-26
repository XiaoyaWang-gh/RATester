package alils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetAppliedMachineGroup_1(t *testing.T) {
	type fields struct {
		Name           string
		InputType      string
		InputDetail    InputDetail
		OutputType     string
		OutputDetail   OutputDetail
		CreateTime     uint32
		LastModifyTime uint32
		project        *LogProject
	}
	type args struct {
		confName string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantGroupNames []string
		wantErr        bool
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

			c := &LogConfig{
				Name:           tt.fields.Name,
				InputType:      tt.fields.InputType,
				InputDetail:    tt.fields.InputDetail,
				OutputType:     tt.fields.OutputType,
				OutputDetail:   tt.fields.OutputDetail,
				CreateTime:     tt.fields.CreateTime,
				LastModifyTime: tt.fields.LastModifyTime,
				project:        tt.fields.project,
			}
			gotGroupNames, err := c.GetAppliedMachineGroup(tt.args.confName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAppliedMachineGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGroupNames, tt.wantGroupNames) {
				t.Errorf("GetAppliedMachineGroup() gotGroupNames = %v, want %v", gotGroupNames, tt.wantGroupNames)
			}
		})
	}
}
