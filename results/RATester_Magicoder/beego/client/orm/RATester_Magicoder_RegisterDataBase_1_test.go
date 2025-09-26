package orm

import (
	"fmt"
	"testing"
)

func TestRegisterDataBase_1(t *testing.T) {
	type args struct {
		aliasName  string
		driverName string
		dataSource string
		params     []DBOption
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				aliasName:  "test",
				driverName: "mysql",
				dataSource: "root:password@tcp(127.0.0.1:3306)/test",
				params:     []DBOption{},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				aliasName:  "test",
				driverName: "mysql",
				dataSource: "root:password@tcp(127.0.0.1:3306)/test",
				params:     []DBOption{},
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

			if err := RegisterDataBase(tt.args.aliasName, tt.args.driverName, tt.args.dataSource, tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("RegisterDataBase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
