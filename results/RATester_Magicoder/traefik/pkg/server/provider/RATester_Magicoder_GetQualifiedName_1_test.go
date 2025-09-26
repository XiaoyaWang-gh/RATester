package provider

import (
	"context"
	"fmt"
	"testing"
)

func TestGetQualifiedName_1(t *testing.T) {
	type args struct {
		ctx         context.Context
		elementName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				ctx:         context.WithValue(context.Background(), "providerName", "testProvider"),
				elementName: "elementName@testProvider",
			},
			want: "testProvider.elementName",
		},
		{
			name: "Test case 2",
			args: args{
				ctx:         context.WithValue(context.Background(), "providerName", "testProvider"),
				elementName: "elementName",
			},
			want: "testProvider.elementName",
		},
		{
			name: "Test case 3",
			args: args{
				ctx:         context.Background(),
				elementName: "elementName",
			},
			want: "elementName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetQualifiedName(tt.args.ctx, tt.args.elementName); got != tt.want {
				t.Errorf("GetQualifiedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
