package echo

import (
	"fmt"
	"testing"
	"time"
)

func TestMustDurations_1(t *testing.T) {
	t.Parallel()

	type args struct {
		sourceParam string
		dest        *[]time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				sourceParam: "sourceParam",
				dest:        &[]time.Duration{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			b := &ValueBinder{}
			if err := b.MustDurations(tt.args.sourceParam, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("MustDurations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
