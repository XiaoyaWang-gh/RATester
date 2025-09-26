package metrics

import (
	"fmt"
	"sync"
	"testing"
)

func TestSetDynamicConfig_1(t *testing.T) {
	type fields struct {
		vectors         []vector
		mtx             sync.Mutex
		dynamicConfig   *dynamicConfig
		deletedEP       []string
		deletedRouters  []string
		deletedServices []string
		deletedURLs     map[string][]string
	}
	type args struct {
		dynamicConfig *dynamicConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			ps := &prometheusState{
				vectors:         tt.fields.vectors,
				mtx:             tt.fields.mtx,
				dynamicConfig:   tt.fields.dynamicConfig,
				deletedEP:       tt.fields.deletedEP,
				deletedRouters:  tt.fields.deletedRouters,
				deletedServices: tt.fields.deletedServices,
				deletedURLs:     tt.fields.deletedURLs,
			}
			ps.SetDynamicConfig(tt.args.dynamicConfig)
		})
	}
}
