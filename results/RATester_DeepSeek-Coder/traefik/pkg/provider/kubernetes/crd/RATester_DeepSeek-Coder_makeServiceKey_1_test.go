package crd

import (
	"fmt"
	"testing"
)

func TestMakeServiceKey_1(t *testing.T) {
	type args struct {
		rule        string
		ingressName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				rule:        "rule1",
				ingressName: "ingress1",
			},
			want:    "ingress1-7f83b1657f",
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				rule:        "rule2",
				ingressName: "ingress2",
			},
			want:    "ingress2-7f83b1657f",
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				rule:        "rule3",
				ingressName: "ingress3",
			},
			want:    "ingress3-7f83b1657f",
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

			got, err := makeServiceKey(tt.args.rule, tt.args.ingressName)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeServiceKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("makeServiceKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
