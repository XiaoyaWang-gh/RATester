package acme

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGetTokenValue_1(t *testing.T) {
	ctx := context.Background()
	challenge := &ChallengeHTTP{
		httpChallenges: map[string]map[string][]byte{
			"token1": {
				"domain1": []byte("value1"),
				"domain2": []byte("value2"),
			},
			"token2": {
				"domain1": []byte("value3"),
				"domain2": []byte("value4"),
			},
		},
	}

	tests := []struct {
		name   string
		ctx    context.Context
		token  string
		domain string
		want   []byte
	}{
		{
			name:   "success",
			ctx:    ctx,
			token:  "token1",
			domain: "domain1",
			want:   []byte("value1"),
		},
		{
			name:   "token not found",
			ctx:    ctx,
			token:  "token3",
			domain: "domain1",
			want:   nil,
		},
		{
			name:   "domain not found",
			ctx:    ctx,
			token:  "token1",
			domain: "domain3",
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := challenge.getTokenValue(tt.ctx, tt.token, tt.domain)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTokenValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
