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

// UsingMemoryDB environment variable that Identify your application will run using a memory based repository (collection of `maps`) or a real database
func UsingMemoryDB() bool {
	usingMemoryDB, _ := strconv.ParseBool(os.Getenv("USE_MEMORY_DB"))
	return usingMemoryDB
}

// ExecuteAutoMigrate environment variable that when true, the application will run as a Auto migration mode and will create the needed tables
func ExecuteAutoMigrate() bool {
	executeAutoMigrate, _ := strconv.ParseBool(os.Getenv("EXECUTE_AUTOMIGRATE"))
	return executeAutoMigrate
}

// JwtSigningKey environment variable that Jwt Signing Key
func JwtSigningKey() string {
	signingKey := os.Getenv("JWT_SIGNING_KEY")
	if signingKey == "" {
		signingKey = "default"
	}
	return signingKey
}
