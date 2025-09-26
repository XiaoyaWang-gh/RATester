package helpers

import (
	"fmt"
	"testing"
)

func TestAdd_5(t *testing.T) {
	type fields struct {
		Name            string
		Pages           uint64
		PaginatorPages  uint64
		Static          uint64
		ProcessedImages uint64
		Files           uint64
		Aliases         uint64
		Cleaned         uint64
	}
	type args struct {
		counter *uint64
		amount  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint64
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

			s := &ProcessingStats{
				Name:            tt.fields.Name,
				Pages:           tt.fields.Pages,
				PaginatorPages:  tt.fields.PaginatorPages,
				Static:          tt.fields.Static,
				ProcessedImages: tt.fields.ProcessedImages,
				Files:           tt.fields.Files,
				Aliases:         tt.fields.Aliases,
				Cleaned:         tt.fields.Cleaned,
			}
			s.Add(tt.args.counter, tt.args.amount)
			if *tt.args.counter != tt.want {
				t.Errorf("ProcessingStats.Add() = %v, want %v", *tt.args.counter, tt.want)
			}
		})
	}
}
