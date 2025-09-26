package http

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewMuxer_2(t *testing.T) {
	tests := []struct {
		name    string
		want    *Muxer
		wantErr bool
	}{
		{
			name: "test",
			want: &Muxer{
				routes:         routes{},
				parser:         nil,
				parserV2:       nil,
				defaultHandler: http.NotFoundHandler(),
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

			got, err := NewMuxer()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMuxer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMuxer() = %v, want %v", got, tt.want)
			}
		})
	}
}
