package client

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeout_1(t *testing.T) {
	type fields struct {
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name: "Test case 1",
			fields: fields{
				timeout: 10 * time.Second,
			},
			want: 10 * time.Second,
		},
		{
			name: "Test case 2",
			fields: fields{
				timeout: 20 * time.Second,
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

			r := &Request{
				timeout: tt.fields.timeout,
			}
			if got := r.Timeout(); got != tt.want {
				t.Errorf("Timeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
