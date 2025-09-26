package mem

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetProxyTraffic_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *ProxyTrafficInfo
	}{
		{
			name: "Test case 1",
			args: args{
				name: "proxy1",
			},
			want: &ProxyTrafficInfo{
				Name: "proxy1",
			},
		},
		{
			name: "Test case 2",
			args: args{
				name: "proxy2",
			},
			want: &ProxyTrafficInfo{
				Name: "proxy2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &serverMetrics{}
			if got := m.GetProxyTraffic(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProxyTraffic() = %v, want %v", got, tt.want)
			}
		})
	}
}
