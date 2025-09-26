package page

import (
	"fmt"
	"testing"
)

func TestHasPrev_1(t *testing.T) {
	type fields struct {
		number    int
		Paginator *Paginator
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test case 1",
			fields: fields{
				number: 1,
				Paginator: &Paginator{
					total: 10,
					size:  10,
				},
			},
			want: false,
		},
		{
			name: "Test case 2",
			fields: fields{
				number: 2,
				Paginator: &Paginator{
					total: 10,
					size:  10,
				},
			},
			want: true,
		},
		{
			name: "Test case 3",
			fields: fields{
				number: 10,
				Paginator: &Paginator{
					total: 100,
					size:  10,
				},
			},
			want: true,
		},
		{
			name: "Test case 4",
			fields: fields{
				number: 11,
				Paginator: &Paginator{
					total: 100,
					size:  10,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Pager{
				number:    tt.fields.number,
				Paginator: tt.fields.Paginator,
			}
			if got := p.HasPrev(); got != tt.want {
				t.Errorf("Pager.HasPrev() = %v, want %v", got, tt.want)
			}
		})
	}
}
