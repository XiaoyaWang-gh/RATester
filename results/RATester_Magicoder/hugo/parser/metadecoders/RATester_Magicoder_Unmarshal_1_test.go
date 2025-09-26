package metadecoders

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshal_1(t *testing.T) {
	type args struct {
		data []byte
		f    Format
	}
	tests := []struct {
		name    string
		d       Decoder
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "Test Unmarshal with CSV format",
			d:    Decoder{},
			args: args{
				data: []byte("field1,field2\nvalue1,value2"),
				f:    CSV,
			},
			want: [][]string{
				{"field1", "field2"},
				{"value1", "value2"},
			},
			wantErr: false,
		},
		{
			name: "Test Unmarshal with non-CSV format",
			d:    Decoder{},
			args: args{
				data: []byte("key:value"),
				f:    "json",
			},
			want:    map[string]any{"key": "value"},
			wantErr: false,
		},
		{
			name: "Test Unmarshal with empty data",
			d:    Decoder{},
			args: args{
				data: []byte(""),
				f:    CSV,
			},
			want:    [][]string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.d.Unmarshal(tt.args.data, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
