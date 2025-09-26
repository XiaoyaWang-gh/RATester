package captcha

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestVerifyReq_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	captcha := &Captcha{
		FieldIDName:      "id",
		FieldCaptchaName: "captcha",
	}

	req, err := http.NewRequest("POST", "/test", strings.NewReader("id=123&captcha=abc"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	result := captcha.VerifyReq(req)
	if !result {
		t.Error("Expected VerifyReq to return true, but got false")
	}
}
