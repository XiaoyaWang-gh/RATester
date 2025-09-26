package port

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalFromName_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *nameBuilder
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				name: "test-name",
			},
			want: &nameBuilder{
				name: "test-name",
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				name: "test-name-100-200",
			},
			want: &nameBuilder{
				name:          "test-name",
				rangePortFrom: 100,
				rangePortTo:   200,
			},
			wantErr: false,
		},
		{
			name: "test case 3",
			args: args{
				name: "test-name-invalid-port",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test case 4",
			args: args{
				name: "test-name-invalid-port-200",
			},
			want:    nil,
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

			got, err := unmarshalFromName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalFromName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshalFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}
