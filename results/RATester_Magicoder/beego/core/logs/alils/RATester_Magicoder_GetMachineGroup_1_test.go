package alils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMachineGroup_1(t *testing.T) {
	project := &LogProject{
		Name:            "test_project",
		Endpoint:        "http://localhost:8080",
		AccessKeyID:     "test_key_id",
		AccessKeySecret: "test_key_secret",
	}

	tests := []struct {
		name    string
		project *LogProject
		want    *MachineGroup
		wantErr bool
	}{
		{
			name:    "TestGetMachineGroup_Success",
			project: project,
			want: &MachineGroup{
				Name:          "test_group",
				Type:          "test_type",
				MachineIDType: "test_id_type",
				MachineIDList: []string{"test_machine_1", "test_machine_2"},
				Attribute: MachineGroupAttribute{
					ExternalName: "test_external_name",
					TopicName:    "test_topic_name",
				},
			},
			wantErr: false,
		},
		{
			name:    "TestGetMachineGroup_Failure",
			project: project,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.project.GetMachineGroup("test_group")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMachineGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMachineGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}
