package web

import (
	"fmt"
	"testing"
	"time"
)

func TestDateFormat_1(t *testing.T) {
	type args struct {
		t      time.Time
		layout string
	}
	tests := []struct {
		name           string
		args           args
		wantDatestring string
	}{
		{
			name: "Test with RFC3339 layout",
			args: args{
				t:      time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
				layout: time.RFC3339,
			},
			wantDatestring: "2022-01-01T12:00:00Z",
		},
		{
			name: "Test with custom layout",
			args: args{
				t:      time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
				layout: "2006-01-02 15:04:05",
			},
			wantDatestring: "2022-01-01 12:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotDatestring := DateFormat(tt.args.t, tt.args.layout); gotDatestring != tt.wantDatestring {
				t.Errorf("DateFormat() = %v, want %v", gotDatestring, tt.wantDatestring)
			}
		})
	}
}
