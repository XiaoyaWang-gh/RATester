package plugin

import (
	"fmt"
	"net/http"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestName_1(t *testing.T) {
	type fields struct {
		options v1.HTTPPluginOptions
		url     string
		client  *http.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestName",
			fields: fields{
				options: v1.HTTPPluginOptions{
					Name: "test",
				},
				url:    "http://localhost",
				client: &http.Client{},
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &httpPlugin{
				options: tt.fields.options,
				url:     tt.fields.url,
				client:  tt.fields.client,
			}
			if got := p.Name(); got != tt.want {
				t.Errorf("httpPlugin.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
