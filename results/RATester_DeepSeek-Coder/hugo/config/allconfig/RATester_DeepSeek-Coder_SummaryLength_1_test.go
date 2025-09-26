package allconfig

import (
	"fmt"
	"testing"
)

func TestSummaryLength_1(t *testing.T) {
	type fields struct {
		config *Config
	}
	tests := []struct {
		name   string
		fields fields
		want   int
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
				config: tt.fields.config,
			}
			if got := c.SummaryLength(); got != tt.want {
				t.Errorf("SummaryLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
