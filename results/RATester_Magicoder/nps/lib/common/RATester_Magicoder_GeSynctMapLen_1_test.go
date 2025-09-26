package common

import (
	"fmt"
	"sync"
	"testing"
)

func TestGeSynctMapLen_1(t *testing.T) {
	tests := []struct {
		name string
		m    sync.Map
		want int
	}{
		{
			name: "Test case 1",
			m:    sync.Map{},
			want: 0,
		},
		{
			name: "Test case 2",
			m: func() sync.Map {
				m := sync.Map{}
				m.Store("key1", "value1")
				m.Store("key2", "value2")
				return m
			}(),
			want: 2,
		},
		{
			name: "Test case 3",
			m: func() sync.Map {
				m := sync.Map{}
				m.Store("key1", "value1")
				m.Store("key2", "value2")
				m.Delete("key1")
				return m
			}(),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GeSynctMapLen(tt.m); got != tt.want {
				t.Errorf("GeSynctMapLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
