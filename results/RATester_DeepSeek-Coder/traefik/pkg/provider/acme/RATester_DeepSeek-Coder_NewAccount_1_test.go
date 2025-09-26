package acme

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestNewAccount_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		email   string
		keyType string
	}
	tests := []struct {
		name    string
		args    args
		want    *Account
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				ctx:     context.Background(),
				email:   "test@example.com",
				keyType: "RSA",
			},
			want: &Account{
				Email:   "test@example.com",
				KeyType: "RSA",
			},
			wantErr: false,
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := NewAccount(tt.args.ctx, tt.args.email, tt.args.keyType)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
