package alils

import (
	"fmt"
	"testing"
)

func TestCreateMachineGroup_1(t *testing.T) {
	type args struct {
		m *MachineGroup
	}
	tests := []struct {
		name    string
		p       *LogProject
		args    args
		wantErr bool
	}{
		{
			name: "TestCreateMachineGroup",
			p: &LogProject{
				Name:            "test",
				Endpoint:        "http://localhost:8080",
				AccessKeyID:     "test",
				AccessKeySecret: "test",
			},
			args: args{
				m: &MachineGroup{
					Name:          "test",
					Type:          "test",
					MachineIDType: "test",
					MachineIDList: []string{"test"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tt.p.CreateMachineGroup(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogProject.CreateMachineGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
