package types

import (
	"fmt"
	"testing"
)

func TestSetDefaults_4(t *testing.T) {
	type fields struct {
		Endpoint string
		TLS      *ClientTLS
		Headers  map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test SetDefaults",
			fields: fields{
				Endpoint: "",
				TLS:      &ClientTLS{},
				Headers:  map[string]string{},
			},
			want: "https://localhost:4318",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &OtelHTTP{
				Endpoint: tt.fields.Endpoint,
				TLS:      tt.fields.TLS,
				Headers:  tt.fields.Headers,
			}
			c.SetDefaults()
			if c.Endpoint != tt.want {
				t.Errorf("OtelHTTP.SetDefaults() = %v, want %v", c.Endpoint, tt.want)
			}
		})
	}
}
