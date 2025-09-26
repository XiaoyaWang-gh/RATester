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
		// TODO: Add test cases.
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
				t.Errorf("serverMetrics.GetProxyTraffic() = %v, want %v", got, tt.want)
			}
		})
	}
}
