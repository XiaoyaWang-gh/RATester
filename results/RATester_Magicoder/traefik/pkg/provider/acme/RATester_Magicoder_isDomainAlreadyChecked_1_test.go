package acme

import (
	"fmt"
	"testing"
)

func TestIsDomainAlreadyChecked_1(t *testing.T) {
	type args struct {
		domainToCheck   string
		existentDomains []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Context_0",
			args: args{
				domainToCheck:   "example.com",
				existentDomains: []string{"example.com", "sub.example.com"},
			},
			want: true,
		},
		{
			name: "Context_1",
			args: args{
				domainToCheck:   "example.com",
				existentDomains: []string{"sub.example.com"},
			},
			want: false,
		},
		{
			name: "Context_2",
			args: args{
				domainToCheck:   "sub.example.com",
				existentDomains: []string{"example.com", "sub.example.com"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isDomainAlreadyChecked(tt.args.domainToCheck, tt.args.existentDomains); got != tt.want {
				t.Errorf("isDomainAlreadyChecked() = %v, want %v", got, tt.want)
			}
		})
	}
}
