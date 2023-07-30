package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var secret = viper.GetString("oauth.sessionCookie.secret")

type tokenClaims struct {
	jwt.StandardClaims
	User
}

type User struct {
	ExternalUser
	UserId string `json:"userId"`
}

func GenerateTokenForGuest(sessionId string) (string, error) {
	return GenerateToken(&User{
		ExternalUser: ExternalUser{},
		UserId:       "",
	})
}

func GenerateToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId": user.UserId,
	})

	return token.SignedString([]byte(secret))
}

func ExtractUserAndTokenFromStringToken(tokenString string) (*User, *jwt.Token, error) {
	var claims tokenClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("invalid signing method, expected HMAC, got: %s", token.Method.Alg())
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, nil, err
	}

	return &claims.User, token, nil
}
