package server

import (
	"fmt"
	"testing"
)

func TestStop_3(t *testing.T) {
	testCases := []struct {
		name        string
		entryPoints UDPEntryPoints
	}{
		{
			name:        "empty entry points",
			entryPoints: UDPEntryPoints{},
		},
		{
			name: "single entry point",
			entryPoints: UDPEntryPoints{
				"entryPoint1": &UDPEntryPoint{},
			},
		},
		{
			name: "multiple entry points",
			entryPoints: UDPEntryPoints{
				"entryPoint1": &UDPEntryPoint{},
				"entryPoint2": &UDPEntryPoint{},
				"entryPoint3": &UDPEntryPoint{},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			test.entryPoints.Stop()
		})
	}
}
