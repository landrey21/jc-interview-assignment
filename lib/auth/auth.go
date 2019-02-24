package auth

import (
	"crypto/sha512"
	"encoding/base64"
)

// hash and encode password
func HashEncode(password string) string {
	//log.Printf("hashEncode() %s", password) // DO NOT PRINT THIS EXCEPT TO TEST!!

	sha := sha512.New()
	sha.Write([]byte(password))
	encodedPassword := base64.StdEncoding.EncodeToString(sha.Sum(nil))

	return encodedPassword
}
