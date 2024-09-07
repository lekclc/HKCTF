package logic

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var jwtKey = []byte(viper.GetString("key.jwt"))

// Claim is a struct that will be encoded to a JWT.
type Claim struct {
	Username string `json:"username"`
	Auth     bool   `json:"auth"`
	jwt.RegisteredClaims
}

func Jwt_get(username string, auth bool) string {
	claims := &Claim{
		Username: username,
		Auth:     auth,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, _ := token.SignedString(jwtKey)
	return tokenstring
}
func Jwt_parse(tokenstring string) (*Claim, error) {
	claims := &Claim{}
	token, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return claims, nil
	}
	return nil, err
}
func Jwt_update(tokenstring string) string {
	//当token有效时长小于12小时,刷新token
	claims, err := Jwt_parse(tokenstring)
	if err != nil {
		return ""
	}
	if time.Now().Add(time.Hour * 12).After(claims.ExpiresAt.Time) {
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenstring, _ := token.SignedString(jwtKey)
		return tokenstring
	}
	return ""
}
func Jwt_check(r *gin.Context, token string) int {
	if token != "" {
		claims, err := Jwt_parse(token)
		if err != nil {
			Res_msg(r, 200, 0, "token error")
			return 0
		}
		token_ := Jwt_update(token)
		if claims.Auth {
			Res_msg(r, 200, 2, "admin login success", token_)
			return 2
		}
		Res_msg(r, 200, 1, "user login success", token_)
		return 1
	}
	return 0
}

func Jwt_Info(r *gin.Context) (string, bool) {
	token := r.GetHeader("Token")
	if token == "" {
		Res_msg(r, 200, 0, "no token")
		return "", false
	}
	claims, err := Jwt_parse(token)
	if err != nil {
		Res_msg(r, 200, 0, "parse token failed")
		return "", false
	}
	if time.Now().After(claims.ExpiresAt.Time) {
		Res_msg(r, 200, 0, "token expired")
		return "", false
	}
	return claims.Username, claims.Auth
}
