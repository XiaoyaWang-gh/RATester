package server

import (
	"fmt"
	"testing"
	"time"
)

func TestflowSession_1(t *testing.T) {
	type args struct {
		m time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				m: 1 * time.Second,
			},
		},
		{
			name: "Test case 2",
			args: args{
				m: 2 * time.Second,
			},
		},
		{
			name: "Test case 3",
			args: args{
				m: 3 * time.Second,
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

			flowSession(tt.args.m)
		})
	}
}
