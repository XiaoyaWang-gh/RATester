package fiber

import (
	"fmt"
	"testing"
)

func TestGetOffer_1(t *testing.T) {
	type args struct {
		header     []byte
		isAccepted func(spec, offer string, specParams headerParams) bool
		offers     []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getOffer(tt.args.header, tt.args.isAccepted, tt.args.offers...); got != tt.want {
				t.Errorf("getOffer() = %v, want %v", got, tt.want)
			}
		})
	}
}
