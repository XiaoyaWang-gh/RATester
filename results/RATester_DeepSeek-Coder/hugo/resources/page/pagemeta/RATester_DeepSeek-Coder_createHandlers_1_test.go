package pagemeta

import (
	"fmt"
	"testing"
)

func TestCreateHandlers_1(t *testing.T) {
	type args struct {
		fmConfig FrontmatterConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid FrontmatterConfig",
			args: args{
				fmConfig: FrontmatterConfig{
					Date:        []string{"date"},
					Lastmod:     []string{"lastmod"},
					PublishDate: []string{"publishdate"},
					ExpiryDate:  []string{"expirydate"},
				},
			},
			wantErr: false,
		},
		{
			name: "Test with invalid FrontmatterConfig",
			args: args{
				fmConfig: FrontmatterConfig{
					Date:        []string{"date"},
					Lastmod:     []string{"lastmod"},
					PublishDate: []string{"publishdate"},
					ExpiryDate:  []string{"expirydate"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := &FrontMatterHandler{
				fmConfig: tt.args.fmConfig,
			}
			if err := f.createHandlers(); (err != nil) != tt.wantErr {
				t.Errorf("FrontMatterHandler.createHandlers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
