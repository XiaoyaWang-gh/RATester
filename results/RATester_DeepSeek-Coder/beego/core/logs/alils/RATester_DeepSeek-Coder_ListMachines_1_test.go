package alils

import (
	"fmt"
	"testing"
)

func TestListMachines_1(t *testing.T) {
	type args struct {
		groupName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				groupName: "test_group",
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				groupName: "non_exist_group",
			},
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

			m := &MachineGroup{
				Name: tt.args.groupName,
			}
			_, _, err := m.ListMachines()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMachines() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
