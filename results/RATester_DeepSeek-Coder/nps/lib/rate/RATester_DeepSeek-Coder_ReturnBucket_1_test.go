package rate

import (
	"fmt"
	"testing"
)

func TestReturnBucket_1(t *testing.T) {
	type fields struct {
		bucketSize        int64
		bucketSurplusSize int64
		bucketAddSize     int64
		stopChan          chan bool
		NowRate           int64
	}
	type args struct {
		size int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
				bucketSize:        tt.fields.bucketSize,
				bucketSurplusSize: tt.fields.bucketSurplusSize,
				bucketAddSize:     tt.fields.bucketAddSize,
				stopChan:          tt.fields.stopChan,
				NowRate:           tt.fields.NowRate,
			}
			s.ReturnBucket(tt.args.size)
		})
	}
}
