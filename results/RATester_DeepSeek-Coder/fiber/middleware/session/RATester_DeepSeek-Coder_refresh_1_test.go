package session

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
)

func TestRefresh_1(t *testing.T) {
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
			s.refresh()
			if s.id != tt.want {
				t.Errorf("Session.refresh() = %v, want %v", s.id, tt.want)
			}
			if s.fresh != true {
				t.Errorf("Session.refresh() = %v, want %v", s.fresh, true)
			}
		})
	}
}
