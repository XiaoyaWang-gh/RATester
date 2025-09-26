package limiter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalMsg_1(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    item
		wantErr bool
	}{
		{
			name:  "Test case 1",
			input: []byte(`{"currHits": 1, "prevHits": 2, "exp": 3}`),
			want: item{
				currHits: 1,
				prevHits: 2,
				exp:      3,
			},
			wantErr: false,
		},
		{
			name:  "Test case 2",
			input: []byte(`{"currHits": 4, "prevHits": 5, "exp": 6}`),
			want: item{
				currHits: 4,
				prevHits: 5,
				exp:      6,
			},
			wantErr: false,
		},
		{
			name:    "Test case 3",
			input:   []byte(`{"currHits": 7, "prevHits": 8, "exp": 9}`),
			want:    item{},
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

			got := &item{}
			_, err := got.UnmarshalMsg(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("UnmarshalMsg() got = %v, want %v", *got, tt.want)
			}
		})
	}
}
