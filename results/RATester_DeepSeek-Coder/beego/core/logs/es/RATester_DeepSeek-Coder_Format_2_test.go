package es

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/logs"
	"github.com/elastic/go-elasticsearch/v6"
)

func TestFormat_2(t *testing.T) {
	type fields struct {
		Client      *elasticsearch.Client
		DSN         string
		Level       int
		formatter   logs.LogFormatter
		Formatter   string
		indexNaming IndexNaming
	}
	type args struct {
		lm *logs.LogMsg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			el := &esLogger{
				Client:      tt.fields.Client,
				DSN:         tt.fields.DSN,
				Level:       tt.fields.Level,
				formatter:   tt.fields.formatter,
				Formatter:   tt.fields.Formatter,
				indexNaming: tt.fields.indexNaming,
			}
			if got := el.Format(tt.args.lm); got != tt.want {
				t.Errorf("esLogger.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
