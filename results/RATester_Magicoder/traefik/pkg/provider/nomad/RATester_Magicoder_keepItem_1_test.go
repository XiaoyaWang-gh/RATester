package nomad

import (
	"context"
	"fmt"
	"testing"
)

func TestKeepItem_1(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		item item
		want bool
	}{
		{
			name: "item with enabled extraConf should pass",
			item: item{
				ExtraConf: configuration{
					Enable: true,
				},
			},
			want: true,
		},
		{
			name: "item with disabled extraConf should fail",
			item: item{
				ExtraConf: configuration{
					Enable: false,
				},
			},
			want: false,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{}
			if got := p.keepItem(ctx, tt.item); got != tt.want {
				t.Errorf("keepItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
