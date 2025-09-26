package fiber

import (
	"fmt"
	"testing"
)

func TestprepareListenData_1(t *testing.T) {
	app := new(App)

	tests := []struct {
		name   string
		addr   string
		isTLS  bool
		cfg    ListenConfig
		expect ListenData
	}{
		{
			name:  "Test case 1",
			addr:  "localhost:8080",
			isTLS: false,
			cfg: ListenConfig{
				ListenerNetwork: NetworkTCP4,
			},
			expect: ListenData{
				Host: "localhost",
				Port: "8080",
				TLS:  false,
			},
		},
		{
			name:  "Test case 2",
			addr:  "",
			isTLS: true,
			cfg: ListenConfig{
				ListenerNetwork: NetworkTCP6,
			},
			expect: ListenData{
				Host: "[::1]",
				Port: "",
				TLS:  true,
			},
		},
		{
			name:  "Test case 3",
			addr:  ":8080",
			isTLS: false,
			cfg: ListenConfig{
				ListenerNetwork: NetworkTCP4,
			},
			expect: ListenData{
				Host: globalIpv4Addr,
				Port: "8080",
				TLS:  false,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := app.prepareListenData(test.addr, test.isTLS, test.cfg)
			if result != test.expect {
				t.Errorf("Expected %v, but got %v", test.expect, result)
			}
		})
	}
}
