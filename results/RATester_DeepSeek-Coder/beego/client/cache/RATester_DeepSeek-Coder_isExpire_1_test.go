package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestIsExpire_1(t *testing.T) {
	type fields struct {
		val         interface{}
		createdTime time.Time
		lifespan    time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test case 1",
			fields: fields{
				val:         "test",
				createdTime: time.Now(),
				lifespan:    1 * time.Second,
			},
			want: true,
		},
		{
			name: "Test case 2",
			fields: fields{
				val:         "test",
				createdTime: time.Now(),
				lifespan:    0,
			},
			want: false,
		},
		{
			name: "Test case 3",
			fields: fields{
				val:         "test",
				createdTime: time.Now().Add(-1 * time.Second),
				lifespan:    1 * time.Second,
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

			mi := &MemoryItem{
				val:         tt.fields.val,
				createdTime: tt.fields.createdTime,
				lifespan:    tt.fields.lifespan,
			}
			if got := mi.isExpire(); got != tt.want {
				t.Errorf("MemoryItem.isExpire() = %v, want %v", got, tt.want)
			}
		})
	}
}
