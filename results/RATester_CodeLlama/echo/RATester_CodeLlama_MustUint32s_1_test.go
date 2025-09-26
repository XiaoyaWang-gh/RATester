package echo

import (
	"fmt"
	"testing"
)

func TestMustUint32s_1(t *testing.T) {
	t.Parallel()

	type args struct {
		sourceParam string
		dest        *[]uint32
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
				dest:        &[]uint32{},
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
			if err := b.MustUint32s(tt.args.sourceParam, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("MustUint32s() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
