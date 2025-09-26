package output

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetByNames_1(t *testing.T) {
	formats := Formats{
		{Name: "html"},
		{Name: "rss"},
		{Name: "atom"},
	}

	tests := []struct {
		names   []string
		want    Formats
		wantErr bool
	}{
		{
			names: []string{"html", "rss", "atom"},
			want: Formats{
				{Name: "html"},
				{Name: "rss"},
				{Name: "atom"},
			},
			wantErr: false,
		},
		{
			names:   []string{"html", "rss", "invalid"},
			want:    Formats{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("names=%v", tt.names), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := formats.GetByNames(tt.names...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
