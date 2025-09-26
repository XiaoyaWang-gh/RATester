package nathole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestListLocalIPsForNatHole_1(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				max: 1,
			},
			want:    []string{"192.168.1.1"},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				max: 2,
			},
			want:    []string{"192.168.1.1", "192.168.1.2"},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				max: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test case 4",
			args: args{
				max: -1,
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

			got, err := ListLocalIPsForNatHole(tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListLocalIPsForNatHole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListLocalIPsForNatHole() = %v, want %v", got, tt.want)
			}
		})
	}
}
