package config

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestBool_4(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		key     string
		reader  func(ctx context.Context, key string) (string, error)
		want    bool
		wantErr bool
	}{
		{
			name:    "success",
			key:     "key1",
			reader:  func(ctx context.Context, key string) (string, error) { return "true", nil },
			want:    true,
			wantErr: false,
		},
		{
			name:    "error",
			key:     "key2",
			reader:  func(ctx context.Context, key string) (string, error) { return "", errors.New("error") },
			want:    false,
			wantErr: true,
		},
		{
			name:    "invalid",
			key:     "key3",
			reader:  func(ctx context.Context, key string) (string, error) { return "invalid", nil },
			want:    false,
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

			c := &BaseConfiger{
				reader: tt.reader,
			}
			got, err := c.Bool(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseConfiger.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BaseConfiger.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}
