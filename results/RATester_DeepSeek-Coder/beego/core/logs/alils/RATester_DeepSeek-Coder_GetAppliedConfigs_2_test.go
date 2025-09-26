package alils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetAppliedConfigs_2(t *testing.T) {
	type fields struct {
		Name           string
		Type           string
		MachineIDType  string
		MachineIDList  []string
		Attribute      MachineGroupAttribute
		CreateTime     uint32
		LastModifyTime uint32
		project        *LogProject
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
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

			m := &MachineGroup{
				Name:           tt.fields.Name,
				Type:           tt.fields.Type,
				MachineIDType:  tt.fields.MachineIDType,
				MachineIDList:  tt.fields.MachineIDList,
				Attribute:      tt.fields.Attribute,
				CreateTime:     tt.fields.CreateTime,
				LastModifyTime: tt.fields.LastModifyTime,
				project:        tt.fields.project,
			}
			got, err := m.GetAppliedConfigs()
			if (err != nil) != tt.wantErr {
				t.Errorf("MachineGroup.GetAppliedConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MachineGroup.GetAppliedConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}
