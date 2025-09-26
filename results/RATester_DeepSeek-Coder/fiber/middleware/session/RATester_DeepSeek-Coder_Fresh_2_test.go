package session

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
)

func TestFresh_2(t *testing.T) {
	t.Parallel()

	type fields struct {
		ctx        fiber.Ctx
		config     *Store
		data       *data
		byteBuffer *bytes.Buffer
		id         string
		exp        time.Duration
		mu         sync.RWMutex
		fresh      bool
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test Fresh() when fresh is true",
			fields: fields{
				fresh: true,
			},
			want: true,
		},
		{
			name: "Test Fresh() when fresh is false",
			fields: fields{
				fresh: false,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &Session{
				ctx:        tt.fields.ctx,
				config:     tt.fields.config,
				data:       tt.fields.data,
				byteBuffer: tt.fields.byteBuffer,
				id:         tt.fields.id,
				exp:        tt.fields.exp,
				mu:         tt.fields.mu,
				fresh:      tt.fields.fresh,
			}
			if got := s.Fresh(); got != tt.want {
				t.Errorf("Session.Fresh() = %v, want %v", got, tt.want)
			}
		})
	}
}
