package jwt

import (
	"fmt"
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/evandroferreiras/gopher-city-bank/app/common/envvar"
	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
)

func getJWTSecret() []byte {
	return []byte(envvar.JwtSigningKey())
}

// GenerateJWT creates a JWT and embed with the given Id
func GenerateJWT(id string) string {
	var JWTSecret = getJWTSecret()

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}

// GetIDFromJwt returns id inside a jwtToken
func GetIDFromJwt(jwtToken string) (string, error) {
	token, err := jwtGo.Parse(jwtToken, func(token *jwtGo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return getJWTSecret(), nil
	})
	if err != nil {
		return "", errors.Wrap(err, "parse jwt")
	}
	if claims, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
		ID := claims["id"].(string)
		return ID, nil
	}
	return "", errors.New("cant find id on jwt")
}
