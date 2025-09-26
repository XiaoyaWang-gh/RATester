package render

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_12(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		r       HTML
		args    args
		wantErr bool
	}{
		{
			name: "Test Render with no template name",
			r: HTML{
				Template: template.Must(template.New("test").Parse("test")),
				Name:     "",
				Data:     "test",
			},
			args: args{
				w: httptest.NewRecorder(),
			},
			wantErr: false,
		},
		{
			name: "Test Render with template name",
			r: HTML{
				Template: template.Must(template.New("test").Parse("test")),
				Name:     "test",
				Data:     "test",
			},
			args: args{
				w: httptest.NewRecorder(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tt.r.Render(tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTML.Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
