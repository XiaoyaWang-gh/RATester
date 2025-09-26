package docker

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetStringMultipleStrict_1(t *testing.T) {
	type args struct {
		labels     map[string]string
		labelNames []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				labels: map[string]string{
					"label1": "value1",
					"label2": "value2",
				},
				labelNames: []string{"label1", "label2"},
			},
			want: map[string]string{
				"label1": "value1",
				"label2": "value2",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				labels: map[string]string{
					"label1": "value1",
				},
				labelNames: []string{"label1", "label2"},
			},
			want:    nil,
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

			got, err := getStringMultipleStrict(tt.args.labels, tt.args.labelNames...)
			if (err != nil) != tt.wantErr {
				t.Errorf("getStringMultipleStrict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStringMultipleStrict() = %v, want %v", got, tt.want)
			}
		})
	}
}
