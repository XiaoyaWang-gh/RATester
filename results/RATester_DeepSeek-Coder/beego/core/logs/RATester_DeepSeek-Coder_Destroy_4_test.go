package logs

import (
	"fmt"
	"testing"
)

func TestDestroy_4(t *testing.T) {
	type fields struct {
		AuthorName  string
		Title       string
		WebhookURL  string
		RedirectURL string
		ImageURL    string
		Level       int
		formatter   LogFormatter
		Formatter   string
	}
	tests := []struct {
		name   string
		fields fields
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

			s := &JLWriter{
				AuthorName:  tt.fields.AuthorName,
				Title:       tt.fields.Title,
				WebhookURL:  tt.fields.WebhookURL,
				RedirectURL: tt.fields.RedirectURL,
				ImageURL:    tt.fields.ImageURL,
				Level:       tt.fields.Level,
				formatter:   tt.fields.formatter,
				Formatter:   tt.fields.Formatter,
			}
			s.Destroy()
		})
	}
}
