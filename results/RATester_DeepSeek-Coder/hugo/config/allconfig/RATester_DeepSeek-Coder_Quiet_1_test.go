package allconfig

import (
	"fmt"
	"testing"
)

func TestQuiet_1(t *testing.T) {
	type fields struct {
		m *Configs
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := ConfigLanguage{
				m: tt.fields.m,
			}
			if got := c.Quiet(); got != tt.want {
				t.Errorf("Quiet() = %v, want %v", got, tt.want)
			}
		})
	}
}
