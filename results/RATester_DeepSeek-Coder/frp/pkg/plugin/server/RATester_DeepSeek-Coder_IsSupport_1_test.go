package plugin

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestIsSupport_1(t *testing.T) {
	type testCase struct {
		name     string
		op       string
		expected bool
		setup    func(t *testing.T) *httpPlugin
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			op:       "op1",
			expected: true,
			setup: func(t *testing.T) *httpPlugin {
				return &httpPlugin{
					options: v1.HTTPPluginOptions{
						Ops: []string{"op1", "op2", "op3"},
					},
				}
			},
		},
		{
			name:     "Test case 2",
			op:       "op4",
			expected: false,
			setup: func(t *testing.T) *httpPlugin {
				return &httpPlugin{
					options: v1.HTTPPluginOptions{
						Ops: []string{"op1", "op2", "op3"},
					},
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := tc.setup(t)
			result := p.IsSupport(tc.op)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
