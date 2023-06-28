package jwt

import (
	"dev_community_server/common"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Payload struct {
	UserId   string
	RoleType common.RoleType
}

type myClaims struct {
	Payload
	jwt.RegisteredClaims
}

type jwtProvider struct {
	secretKey string
}

func NewJwtProvider(secret string) *jwtProvider {
	return &jwtProvider{secretKey: secret}
}

func (jp *jwtProvider) GenerateAccessToken(payload Payload, expiry int) (*string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &myClaims{
		payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(expiry))),
			IssuedAt:  jwt.NewNumericDate(time.Now().Local()),
		},
	})

	token, err := t.SignedString([]byte(jp.secretKey))
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (jp *jwtProvider) Decode(tokenStr string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jp.secretKey), nil
	})

	claims, ok := token.Claims.(*myClaims)

	if err != nil || !token.Valid || !ok {
		return nil, err
	}

	return &claims.Payload, nil
}
