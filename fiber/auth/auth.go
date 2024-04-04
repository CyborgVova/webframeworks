package auth

import (
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var Auth = map[string]string{"vova": "vovapass", "seva": "sevapass"}

func GetBasicAuth(f *fiber.Ctx) (string, string, bool) {
	auth := f.Get("Authorization")
	if len(auth) <= 6 {
		return "", "", false
	}
	b, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		return "", "", false
	}

	cred := string(b)
	index := strings.Index(cred, ":")
	return cred[:index], cred[index+1:], true
}
