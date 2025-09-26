package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMounts_2(t *testing.T) {
	type fields struct {
		mounts []Mount
	}
	tests := []struct {
		name   string
		fields fields
		want   []Mount
	}{
		{
			name: "Test case 1",
			fields: fields{
				mounts: []Mount{
					{
						Source: "scss",
						Target: "assets/bootstrap/scss",
					},
				},
			},
			want: []Mount{
				{
					Source: "scss",
					Target: "assets/bootstrap/scss",
				},
			},
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

			m := &moduleAdapter{
				mounts: tt.fields.mounts,
			}
			if got := m.Mounts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moduleAdapter.Mounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
