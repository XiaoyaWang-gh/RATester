package orm

import (
	"fmt"
	"testing"
	"time"
)

func TestSetDataBaseTZ_1(t *testing.T) {
	type args struct {
		aliasName string
		tz        *time.Location
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				aliasName: "testAlias",
				tz:        time.Local,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				aliasName: "invalidAlias",
				tz:        time.Local,
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

			if err := SetDataBaseTZ(tt.args.aliasName, tt.args.tz); (err != nil) != tt.wantErr {
				t.Errorf("SetDataBaseTZ() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
