package udp

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestShuffle_1(t *testing.T) {
	type args struct {
		values []int
		r      *rand.Rand
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				values: []int{1, 2, 3, 4, 5},
				r:      rand.New(rand.NewSource(time.Now().UnixNano())),
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := shuffle(tt.args.values, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
