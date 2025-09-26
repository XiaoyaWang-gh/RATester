package hugo

import (
	"fmt"
	"testing"
)

func TestIsServer_1(t *testing.T) {
	type fields struct {
		CommitHash  string
		BuildDate   string
		Environment string
		GoVersion   string
		conf        ConfigProvider
		deps        []*Dependency
		Context     Context
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
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

			i := HugoInfo{
				CommitHash:  tt.fields.CommitHash,
				BuildDate:   tt.fields.BuildDate,
				Environment: tt.fields.Environment,
				GoVersion:   tt.fields.GoVersion,
				conf:        tt.fields.conf,
				deps:        tt.fields.deps,
				Context:     tt.fields.Context,
			}
			if got := i.IsServer(); got != tt.want {
				t.Errorf("HugoInfo.IsServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
