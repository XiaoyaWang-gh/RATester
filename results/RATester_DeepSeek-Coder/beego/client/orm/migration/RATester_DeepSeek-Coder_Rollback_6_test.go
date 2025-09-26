package migration

import (
	"fmt"
	"testing"
)

func TestRollback_6(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestRollback_0",
			args: args{
				name: "test",
			},
			wantErr: false,
		},
		{
			name: "TestRollback_1",
			args: args{
				name: "not_exist",
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

			err := Rollback(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
