package acme

import (
	"fmt"
	"testing"
)

func TestCleanChan_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		c    *ChallengeTLSALPN
		args args
	}{
		{
			name: "cleanChan",
			c:    &ChallengeTLSALPN{chans: make(map[string]chan struct{})},
			args: args{key: "testKey"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.c.cleanChan(tt.args.key)
			if _, ok := tt.c.chans[tt.args.key]; ok {
				t.Errorf("cleanChan() = %v, want %v", tt.c.chans[tt.args.key], nil)
			}
		})
	}
}
