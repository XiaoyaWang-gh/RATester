package fiber

import (
	"fmt"
	"testing"
)

func TestAcceptsOffer_1(t *testing.T) {
	type args struct {
		spec         string
		offer        string
		headerParams headerParams
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				spec:         "*",
				offer:        "offer",
				headerParams: headerParams{"key": []byte("value")},
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				spec:         "offer*",
				offer:        "offer",
				headerParams: headerParams{"key": []byte("value")},
			},
			want: true,
		},
		{
			name: "Test case 3",
			args: args{
				spec:         "other",
				offer:        "offer",
				headerParams: headerParams{"key": []byte("value")},
			},
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

			if got := acceptsOffer(tt.args.spec, tt.args.offer, tt.args.headerParams); got != tt.want {
				t.Errorf("acceptsOffer() = %v, want %v", got, tt.want)
			}
		})
	}
}
