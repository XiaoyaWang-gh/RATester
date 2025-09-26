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
			name: "Test with valid context and element name",
			args: args{
				ctx:         context.WithValue(context.Background(), key, "provider"),
				elementName: "element",
			},
			want: "provider.element",
		},
		{
			name: "Test with valid context and element name with @",
			args: args{
				ctx:         context.WithValue(context.Background(), key, "provider"),
				elementName: "element@domain",
			},
			want: "provider.element",
		},
		{
			name: "Test with invalid context and valid element name",
			args: args{
				ctx:         context.Background(),
				elementName: "element",
			},
			want: "element",
		},
		{
			name: "Test with invalid context and invalid element name",
			args: args{
				ctx:         context.Background(),
				elementName: "@domain",
			},
			want: "@domain",
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
