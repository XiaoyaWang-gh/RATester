package proxy

import (
	"fmt"
	"testing"
)

func TestGetName_1(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestGetName",
			fields: fields{
				name: "TestName",
			},
			want: "TestName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			pxy := &BaseProxy{
				name: tt.fields.name,
			}
			if got := pxy.GetName(); got != tt.want {
				t.Errorf("BaseProxy.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
