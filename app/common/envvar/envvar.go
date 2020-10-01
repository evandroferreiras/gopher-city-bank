package envvar

import (
	"os"
	"strconv"
)

// IsLocalEnv environment variable that returns if is running local or not
func IsLocalEnv() bool {
	isLocalEnv, _ := strconv.ParseBool(os.Getenv("LOCAL_ENV"))
	return isLocalEnv
}

// JwtSigningKey environment variable that Jwt Signing Key
func JwtSigningKey() string {
	signingKey := os.Getenv("JWT_SIGNING_KEY")
	if signingKey == "" {
		signingKey = "default"
	}
	return signingKey
}
