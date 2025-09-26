package allconfig

import (
	"fmt"
	"testing"
)

func TestWorkingDir_1(t *testing.T) {
	type fields struct {
		m *Configs
	}
	tests := []struct {
		name   string
		fields fields
		want   string
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
			if got := c.WorkingDir(); got != tt.want {
				t.Errorf("WorkingDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
