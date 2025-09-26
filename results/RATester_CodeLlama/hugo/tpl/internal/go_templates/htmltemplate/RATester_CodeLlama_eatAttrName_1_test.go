package template

import (
	"fmt"
	"testing"
)

func TestEatAttrName_1(t *testing.T) {
	type args struct {
		s []byte
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"test_case_1",
			args{
				[]byte(""),
				0,
			},
			0,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := eatAttrName(tt.args.s, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("eatAttrName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("eatAttrName() = %v, want %v", got, tt.want)
			}
		})
	}
}
