package alils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestListMachineGroup_1(t *testing.T) {
	type args struct {
		offset int
		size   int
	}
	tests := []struct {
		name      string
		p         *LogProject
		args      args
		wantM     []string
		wantTotal int
		wantErr   bool
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

			gotM, gotTotal, err := tt.p.ListMachineGroup(tt.args.offset, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMachineGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("ListMachineGroup() gotM = %v, want %v", gotM, tt.wantM)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListMachineGroup() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
