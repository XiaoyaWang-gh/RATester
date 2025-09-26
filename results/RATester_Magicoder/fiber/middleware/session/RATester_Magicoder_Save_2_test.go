package session

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
)

func TestSave_2(t *testing.T) {
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
		name    string
		fields  fields
		wantErr bool
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
			if err := s.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Session.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
