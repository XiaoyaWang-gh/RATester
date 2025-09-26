package template

import (
	"fmt"
	"testing"
)

func TestCheckCanParse_1(t *testing.T) {
	type args struct {
		t *Template
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				t: &Template{
					nameSpace: &nameSpace{
						escaped: true,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "case2",
			args: args{
				t: &Template{
					nameSpace: &nameSpace{
						escaped: false,
					},
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

			if err := tt.args.t.checkCanParse(); (err != nil) != tt.wantErr {
				t.Errorf("checkCanParse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
