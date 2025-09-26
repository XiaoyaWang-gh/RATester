package snicheck

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNew_26(t *testing.T) {
	type args struct {
		tlsOptionsForHost map[string]string
		next              http.Handler
	}
	tests := []struct {
		name string
		args args
		want *SNICheck
	}{
		{
			name: "Test Case 1",
			args: args{
				tlsOptionsForHost: map[string]string{"example.com": "tlsOptions"},
				next:              http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
			want: &SNICheck{
				next:              http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
				tlsOptionsForHost: map[string]string{"example.com": "tlsOptions"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := New(tt.args.tlsOptionsForHost, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
