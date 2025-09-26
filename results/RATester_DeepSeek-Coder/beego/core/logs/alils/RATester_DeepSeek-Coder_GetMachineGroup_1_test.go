package alils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMachineGroup_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		p       *LogProject
		args    args
		wantM   *MachineGroup
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

			gotM, err := tt.p.GetMachineGroup(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMachineGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("GetMachineGroup() gotM = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}
