package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtWrapper struct {
	SecretKey        string
	Issuer           string
	ExpirationSecond int64
}

type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

//GenerateToken generate a jwt token with claims
func (_self *JwtWrapper) GenerateToken(email string) (string, error) {
	claims := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(_self.ExpirationSecond)).Unix(),
			Issuer:    "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(_self.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

//ValidateToken validates the jwt token
func (_self *JwtWrapper) ValidateToken(signedToken string) (*JwtClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(_self.SecretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("jwt is expired")
	}

	return claims, nil
}
