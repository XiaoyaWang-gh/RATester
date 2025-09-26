package rate

import (
	"fmt"
	"testing"
)

func TestAdd_2(t *testing.T) {
	type args struct {
		size int64
	}
	tests := []struct {
		name string
		s    *Rate
		args args
		want int64
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

			s := &Rate{
				bucketSize:        tt.s.bucketSize,
				bucketSurplusSize: tt.s.bucketSurplusSize,
				bucketAddSize:     tt.s.bucketAddSize,
				stopChan:          tt.s.stopChan,
				NowRate:           tt.s.NowRate,
			}
			s.add(tt.args.size)
			if got := s.bucketSurplusSize; got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
