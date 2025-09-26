package auth

import (
	"fmt"
	"testing"
)

func TestDigestUserParser_1(t *testing.T) {
	tests := []struct {
		name      string
		user      string
		wantUser  string
		wantRealm string
		wantErr   bool
	}{
		{
			name:      "valid user",
			user:      "user:realm:password",
			wantUser:  "user:realm",
			wantRealm: "password",
			wantErr:   false,
		},
		{
			name:      "invalid user",
			user:      "user:realm",
			wantUser:  "",
			wantRealm: "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotUser, gotRealm, err := digestUserParser(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("digestUserParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUser != tt.wantUser {
				t.Errorf("digestUserParser() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
			if gotRealm != tt.wantRealm {
				t.Errorf("digestUserParser() gotRealm = %v, want %v", gotRealm, tt.wantRealm)
			}
		})
	}
}
