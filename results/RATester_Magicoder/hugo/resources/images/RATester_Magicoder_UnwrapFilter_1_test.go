package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestUnwrapFilter_1(t *testing.T) {
	type args struct {
		in gift.Filter
	}
	tests := []struct {
		name string
		args args
		want gift.Filter
	}{
		{
			name: "Test case 1",
			args: args{
				in: &filter{
					Options: filterOpts{
						Version: 1,
						Vals:    "test",
					},
				},
			},
			want: &filter{
				Options: filterOpts{
					Version: 1,
					Vals:    "test",
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				in: &filter{
					Options: filterOpts{
						Version: 2,
						Vals:    123,
					},
				},
			},
			want: &filter{
				Options: filterOpts{
					Version: 2,
					Vals:    123,
				},
			},
		},
		{
			name: "Test case 3",
			args: args{
				in: &filter{
					Options: filterOpts{
						Version: 3,
						Vals:    true,
					},
				},
			},
			want: &filter{
				Options: filterOpts{
					Version: 3,
					Vals:    true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := UnwrapFilter(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
