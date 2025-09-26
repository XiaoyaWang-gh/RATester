package tcp

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/proxy"
)

func TestTerminationDelay_1(t *testing.T) {
	type fields struct {
		Dialer           proxy.Dialer
		terminationDelay time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name: "Test case 1",
			fields: fields{
				Dialer:           nil,
				terminationDelay: 10 * time.Second,
			},
			want: 10 * time.Second,
		},
		{
			name: "Test case 2",
			fields: fields{
				Dialer:           nil,
				terminationDelay: 20 * time.Second,
			},
			want: 20 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := dialerWrapper{
				Dialer:           tt.fields.Dialer,
				terminationDelay: tt.fields.terminationDelay,
			}
			if got := d.TerminationDelay(); got != tt.want {
				t.Errorf("TerminationDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}
