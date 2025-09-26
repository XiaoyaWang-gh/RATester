package captcha

import (
	"fmt"
	"html/template"
	"testing"
	"time"
)

func TestCreateCaptchaHTML_1(t *testing.T) {
	type fields struct {
		store            Storage
		URLPrefix        string
		FieldIDName      string
		FieldCaptchaName string
		StdWidth         int
		StdHeight        int
		ChallengeNums    int
		Expiration       time.Duration
		CachePrefix      string
	}
	tests := []struct {
		name   string
		fields fields
		want   template.HTML
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

			c := &Captcha{
				store:            tt.fields.store,
				URLPrefix:        tt.fields.URLPrefix,
				FieldIDName:      tt.fields.FieldIDName,
				FieldCaptchaName: tt.fields.FieldCaptchaName,
				StdWidth:         tt.fields.StdWidth,
				StdHeight:        tt.fields.StdHeight,
				ChallengeNums:    tt.fields.ChallengeNums,
				Expiration:       tt.fields.Expiration,
				CachePrefix:      tt.fields.CachePrefix,
			}
			if got := c.CreateCaptchaHTML(); got != tt.want {
				t.Errorf("CreateCaptchaHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
