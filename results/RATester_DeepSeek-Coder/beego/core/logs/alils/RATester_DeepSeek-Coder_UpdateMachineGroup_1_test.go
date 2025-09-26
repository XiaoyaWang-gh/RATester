package alils

import (
	"fmt"
	"testing"
)

func TestUpdateMachineGroup_1(t *testing.T) {
	type args struct {
		m *MachineGroup
	}
	tests := []struct {
		name    string
		p       *LogProject
		args    args
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

			p := &LogProject{
				Name:            tt.p.Name,
				Endpoint:        tt.p.Endpoint,
				AccessKeyID:     tt.p.AccessKeyID,
				AccessKeySecret: tt.p.AccessKeySecret,
			}
			if err := p.UpdateMachineGroup(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("LogProject.UpdateMachineGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
