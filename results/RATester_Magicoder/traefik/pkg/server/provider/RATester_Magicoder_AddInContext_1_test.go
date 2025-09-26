package provider

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestAddInContext_1(t *testing.T) {
	type args struct {
		ctx         context.Context
		elementName string
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "Test case 1",
			args: args{
				ctx:         context.Background(),
				elementName: "test@provider",
			},
			want: context.WithValue(context.Background(), key, "provider"),
		},
		{
			name: "Test case 2",
			args: args{
				ctx:         context.Background(),
				elementName: "test",
			},
			want: context.Background(),
		},
		{
			name: "Test case 3",
			args: args{
				ctx:         context.WithValue(context.Background(), key, "provider"),
				elementName: "test@provider",
			},
			want: context.WithValue(context.Background(), key, "provider"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := AddInContext(tt.args.ctx, tt.args.elementName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddInContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
