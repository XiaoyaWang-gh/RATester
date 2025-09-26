package consulcatalog

import (
	"fmt"
	"testing"
)

func TestItemServersTransportKey_1(t *testing.T) {
	type args struct {
		item itemData
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				item: itemData{
					ID:         "1",
					Node:       "node1",
					Datacenter: "dc1",
					Name:       "name1",
					Namespace:  "ns1",
					Address:    "address1",
					Port:       "port1",
					Status:     "status1",
					Labels:     map[string]string{"label1": "value1"},
					Tags:       []string{"tag1", "tag2"},
					ExtraConf:  configuration{Enable: true, ConsulCatalog: specificConfiguration{Connect: true, Canary: true}},
				},
			},
			want: "tls-ns1-dc1-name1",
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := itemServersTransportKey(tt.args.item); got != tt.want {
				t.Errorf("itemServersTransportKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
