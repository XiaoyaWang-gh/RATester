package mem

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetProxiesByTypeAndName_1(t *testing.T) {
	type args struct {
		proxyType string
		proxyName string
	}
	tests := []struct {
		name string
		args args
		want *ProxyStats
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
			if got := m.GetProxiesByTypeAndName(tt.args.proxyType, tt.args.proxyName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serverMetrics.GetProxiesByTypeAndName() = %v, want %v", got, tt.want)
			}
		})
	}
}
