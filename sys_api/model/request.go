package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	ID          uint32 `json:"id"`
	NickName    string `json:"nick_name"`
	AuthorityId uint   `json:"authority_id"`
	jwt.RegisteredClaims
}
