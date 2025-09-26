package externalversions

import (
	"fmt"
	reflect "reflect"
	"testing"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestWithCustomResyncConfig_1(t *testing.T) {
	type args struct {
		resyncConfig map[v1.Object]time.Duration
	}
	tests := []struct {
		name string
		args args
		want *sharedInformerFactory
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

			if got := WithCustomResyncConfig(tt.args.resyncConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCustomResyncConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
