package nathole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestClassifyNATFeature_1(t *testing.T) {
	type args struct {
		addresses []string
		localIPs  []string
	}
	tests := []struct {
		name    string
		args    args
		want    *NatFeature
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				addresses: []string{"192.168.1.100:12345", "192.168.1.100:12346", "192.168.1.100:12347"},
				localIPs:  []string{"192.168.1.100"},
			},
			want: &NatFeature{
				NatType:            "HardNAT",
				Behavior:           "BehaviorPortChanged",
				PortsDifference:    2,
				RegularPortsChange: true,
				PublicNetwork:      true,
			},
			wantErr: false,
		},
		// add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ClassifyNATFeature(tt.args.addresses, tt.args.localIPs)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClassifyNATFeature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClassifyNATFeature() = %v, want %v", got, tt.want)
			}
		})
	}
}
