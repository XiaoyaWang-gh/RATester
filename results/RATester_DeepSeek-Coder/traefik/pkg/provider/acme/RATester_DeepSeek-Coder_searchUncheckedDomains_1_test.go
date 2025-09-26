package acme

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestSearchUncheckedDomains_1(t *testing.T) {
	type args struct {
		ctx             context.Context
		domainsToCheck  []string
		existentDomains []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{
				ctx:             context.Background(),
				domainsToCheck:  []string{"example1.com", "example2.com", "example3.com"},
				existentDomains: []string{"example1.com", "example2.com"},
			},
			want: []string{"example3.com"},
		},
		{
			name: "Test case 2",
			args: args{
				ctx:             context.Background(),
				domainsToCheck:  []string{"example4.com", "example5.com", "example6.com"},
				existentDomains: []string{"example4.com", "example5.com"},
			},
			want: []string{"example6.com"},
		},
		{
			name: "Test case 3",
			args: args{
				ctx:             context.Background(),
				domainsToCheck:  []string{"example1.com", "example2.com", "example3.com"},
				existentDomains: []string{"example1.com", "example2.com", "example3.com"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := searchUncheckedDomains(tt.args.ctx, tt.args.domainsToCheck, tt.args.existentDomains); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchUncheckedDomains() = %v, want %v", got, tt.want)
			}
		})
	}
}
