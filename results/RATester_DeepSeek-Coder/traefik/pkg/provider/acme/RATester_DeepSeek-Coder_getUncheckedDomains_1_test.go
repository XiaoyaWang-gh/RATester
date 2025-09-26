package acme

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGetUncheckedDomains_1(t *testing.T) {
	type args struct {
		ctx            context.Context
		domainsToCheck []string
		tlsStore       string
	}
	tests := []struct {
		name string
		args args
		want []string
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

			p := &Provider{}
			if got := p.getUncheckedDomains(tt.args.ctx, tt.args.domainsToCheck, tt.args.tlsStore); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUncheckedDomains() = %v, want %v", got, tt.want)
			}
		})
	}
}
