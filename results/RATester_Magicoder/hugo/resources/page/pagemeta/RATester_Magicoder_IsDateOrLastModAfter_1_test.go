package pagemeta

import (
	"fmt"
	"testing"
	"time"
)

func TestIsDateOrLastModAfter_1(t *testing.T) {
	type args struct {
		in Dates
	}
	tests := []struct {
		name string
		d    Dates
		args args
		want bool
	}{
		{
			name: "Test case 1",
			d:    Dates{Date: time.Now(), Lastmod: time.Now().Add(time.Hour)},
			args: args{in: Dates{Date: time.Now().Add(time.Hour), Lastmod: time.Now().Add(time.Hour * 2)}},
			want: true,
		},
		{
			name: "Test case 2",
			d:    Dates{Date: time.Now(), Lastmod: time.Now().Add(time.Hour)},
			args: args{in: Dates{Date: time.Now().Add(time.Hour * 2), Lastmod: time.Now().Add(time.Hour * 2)}},
			want: false,
		},
		{
			name: "Test case 3",
			d:    Dates{Date: time.Now(), Lastmod: time.Now().Add(time.Hour)},
			args: args{in: Dates{Date: time.Now(), Lastmod: time.Now().Add(time.Hour * 2)}},
			want: true,
		},
		{
			name: "Test case 4",
			d:    Dates{Date: time.Now(), Lastmod: time.Now().Add(time.Hour)},
			args: args{in: Dates{Date: time.Now(), Lastmod: time.Now()}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.d.IsDateOrLastModAfter(tt.args.in); got != tt.want {
				t.Errorf("IsDateOrLastModAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}
