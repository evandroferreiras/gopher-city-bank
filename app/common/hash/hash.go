package hash

import (
	"crypto/sha256"
	"encoding/base64"
)

// EncryptString encrypts the string passed using sha256
func EncryptString(valueToEncrypt string) string {
	h := sha256.New()
	h.Write([]byte(valueToEncrypt))
	bytes := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(bytes)
}
