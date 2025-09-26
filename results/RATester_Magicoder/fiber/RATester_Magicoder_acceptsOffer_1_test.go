package fiber

import (
	"fmt"
	"testing"
)

func TestacceptsOffer_1(t *testing.T) {
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
			name: "Test Case 1",
			args: args{
				spec:         "*",
				offer:        "offer",
				headerParams: headerParams{},
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				spec:         "offer",
				offer:        "offer",
				headerParams: headerParams{},
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				spec:         "other",
				offer:        "offer",
				headerParams: headerParams{},
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
