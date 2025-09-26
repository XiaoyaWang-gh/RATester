package fiber

import (
	"fmt"
	"testing"
)

func TestSetMatched_1(t *testing.T) {
	type fields struct {
		matched bool
	}
	type args struct {
		matched bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Case 1",
			fields: fields{
				matched: false,
			},
			args: args{
				matched: true,
			},
		},
		{
			name: "Test Case 2",
			fields: fields{
				matched: true,
			},
			args: args{
				matched: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &DefaultCtx{
				matched: tt.fields.matched,
			}
			c.setMatched(tt.args.matched)
			if c.matched != tt.args.matched {
				t.Errorf("Expected matched to be %v, got %v", tt.args.matched, c.matched)
			}
		})
	}
}
