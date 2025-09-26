package config

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestInt64_2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		reader  func(ctx context.Context, key string) (string, error)
		key     string
		want    int64
		wantErr bool
	}{
		{
			name: "success",
			reader: func(ctx context.Context, key string) (string, error) {
				return "123", nil
			},
			key:  "key",
			want: 123,
		},
		{
			name: "error",
			reader: func(ctx context.Context, key string) (string, error) {
				return "", errors.New("reader error")
			},
			key:     "key",
			wantErr: true,
		},
		{
			name: "parse error",
			reader: func(ctx context.Context, key string) (string, error) {
				return "abc", nil
			},
			key:     "key",
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
			got, err := c.Int64(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64() got = %v, want %v", got, tt.want)
			}
		})
	}
}
