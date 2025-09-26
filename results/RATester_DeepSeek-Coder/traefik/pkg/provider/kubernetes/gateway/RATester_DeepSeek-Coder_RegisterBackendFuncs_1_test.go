package gateway

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestRegisterBackendFuncs_1(t *testing.T) {
	testCases := []struct {
		desc        string
		group       string
		kind        string
		builderFunc BuildBackendFunc
	}{
		{
			desc:        "Register a new backend function",
			group:       "testGroup",
			kind:        "testKind",
			builderFunc: func(name, namespace string) (string, *dynamic.Service, error) { return "", nil, nil },
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{}
			p.RegisterBackendFuncs(test.group, test.kind, test.builderFunc)

			groupKindBackendFuncs := p.groupKindBackendFuncs[test.group]
			if groupKindBackendFuncs == nil {
				t.Errorf("Expected group %s to exist in groupKindBackendFuncs", test.group)
			}

			builderFunc := groupKindBackendFuncs[test.kind]
			if builderFunc == nil {
				t.Errorf("Expected kind %s to exist in group %s", test.kind, test.group)
			}
		})
	}
}
