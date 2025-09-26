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
			name: "Test case 1",
			args: args{
				t:      time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
				layout: time.RFC3339,
			},
			wantDatestring: "2022-01-01T12:00:00Z",
		},
		{
			name: "Test case 2",
			args: args{
				t:      time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
				layout: time.RFC1123,
			},
			wantDatestring: "Mon, 01 Jan 2022 12:00 UTC",
		},
		{
			name: "Test case 3",
			args: args{
				t:      time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
				layout: time.RFC822,
			},
			wantDatestring: "01 Jan 22 12:00 UTC",
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
