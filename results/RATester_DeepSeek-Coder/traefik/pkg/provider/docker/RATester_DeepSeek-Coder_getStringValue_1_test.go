package docker

import (
	"fmt"
	"testing"
)

func TestGetStringValue_1(t *testing.T) {
	type args struct {
		labels       map[string]string
		labelName    string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				labels:       map[string]string{"label1": "value1"},
				labelName:    "label1",
				defaultValue: "default",
			},
			want: "value1",
		},
		{
			name: "Test case 2",
			args: args{
				labels:       map[string]string{"label2": ""},
				labelName:    "label2",
				defaultValue: "default",
			},
			want: "default",
		},
		{
			name: "Test case 3",
			args: args{
				labels:       map[string]string{"label3": "value3"},
				labelName:    "label4",
				defaultValue: "default",
			},
			want: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getStringValue(tt.args.labels, tt.args.labelName, tt.args.defaultValue); got != tt.want {
				t.Errorf("getStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
