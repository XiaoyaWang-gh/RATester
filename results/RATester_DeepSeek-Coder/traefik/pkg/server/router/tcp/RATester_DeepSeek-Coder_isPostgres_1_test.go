package tcp

import (
	"bufio"
	"fmt"
	"testing"
)

func TestIsPostgres_1(t *testing.T) {
	type args struct {
		br *bufio.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    bool
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

			got, err := isPostgres(tt.args.br)
			if (err != nil) != tt.wantErr {
				t.Errorf("isPostgres() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isPostgres() = %v, want %v", got, tt.want)
			}
		})
	}
}
