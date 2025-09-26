package echo

import (
	"fmt"
	"testing"
)

func TestMustUint64s_1(t *testing.T) {
	t.Parallel()

	type args struct {
		sourceParam string
		dest        *[]uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				sourceParam: "sourceParam",
				dest:        &[]uint64{},
			},
			wantErr: false,
		},
		{
			name: "failure",
			args: args{
				sourceParam: "sourceParam",
				dest:        nil,
			},
			wantErr: true,
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
			if err := b.MustUint64s(tt.args.sourceParam, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("MustUint64s() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
