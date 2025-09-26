package models

import (
	"fmt"
	"testing"
)

func TestBootstrap_1(t *testing.T) {
	mc := &ModelCache{
		cache:           make(map[string]*ModelInfo),
		cacheByFullName: make(map[string]*ModelInfo),
	}

	tests := []struct {
		name string
		mc   *ModelCache
		want error
	}{
		{
			name: "TestBootstrap",
			mc:   mc,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.mc.Bootstrap()
			if tt.mc.done != (tt.want == nil) {
				t.Errorf("Bootstrap() error = %v, wantErr %v", tt.mc.done, tt.want)
			}
		})
	}
}
