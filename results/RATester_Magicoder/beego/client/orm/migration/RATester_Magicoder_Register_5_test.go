package migration

import (
	"fmt"
	"testing"
)

func TestRegister_5(t *testing.T) {
	type args struct {
		name string
		m    Migrationer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestRegister_Success",
			args: args{
				name: "test",
				m:    &Migration{},
			},
			wantErr: false,
		},
		{
			name: "TestRegister_Fail",
			args: args{
				name: "test",
				m:    &Migration{},
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

			if err := Register(tt.args.name, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
