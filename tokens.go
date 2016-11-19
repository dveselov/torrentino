package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TorrentClaims struct {
	URL  string `json:"URL"`
	Path string `json:"path"`
	jwt.StandardClaims
}

func NewToken(URL string, path string) (string, error) {
	claims := TorrentClaims{
		URL,
		path,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(TokenTTL)).Unix(),
			Issuer:    TokenIssuer,
		},
	}

	token := jwt.NewWithClaims(TokenSigningMethod, claims)
	tokenString, err := token.SignedString(TokenSigningKey)
	return tokenString, err
}

func ParseToken(tokenString string) (*TorrentClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &TorrentClaims{}, func(token *jwt.Token) (interface{}, error) {
		return TokenSigningKey, nil
	})

	if err != nil {
		return &TorrentClaims{}, err
	}

	claims := token.Claims.(*TorrentClaims)
	return claims, err
}
