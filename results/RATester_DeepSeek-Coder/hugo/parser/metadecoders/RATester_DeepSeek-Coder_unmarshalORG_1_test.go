package metadecoders

import (
	"fmt"
	"testing"
)

func TestUnmarshalORG_1(t *testing.T) {
	type args struct {
		data []byte
		v    any
	}
	tests := []struct {
		name    string
		d       Decoder
		args    args
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

			d := Decoder{
				Delimiter:  ',',
				Comment:    '#',
				LazyQuotes: true,
			}
			if err := d.unmarshalORG(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.unmarshalORG() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
