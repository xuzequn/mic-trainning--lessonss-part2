package jwt_op

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"mic-trainning-lessons/internal"
	"mic-trainning-lessons/log"
	"time"
)

const (
	TokenExpried     = "Token已过期"
	TokenNotValidYet = "Token不再有效"
	TokenMalformed   = "Token非法"
	TokenInvalid     = "Token无效"
)

type CustomClaims struct {
	jwt.StandardClaims
	Id          int32
	NikeName    string
	AuthorityId int32 //  权限认证
}

type JWT struct {
	SiginKey []byte
}

func NewJWT() *JWT {
	return &JWT{SiginKey: []byte(internal.AppConf.JWTConfig.SingingKey)}
}

// 生成token
func (j *JWT) GenerateJWT(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(j.SiginKey)
	if err != nil {
		log.Logger.Error("生成JWT错误:" + err.Error())
		return "", err
	}
	return tokenStr, nil
}

// 解析token
func (j *JWT) ParseToken(tokenStr string) (*CustomClaims, error) {
	//  解析token
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SiginKey, nil
	})
	//问题处理
	if err != nil {
		if result, ok := err.(jwt.ValidationError); ok {
			if result.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(TokenMalformed)
			} else if result.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(TokenExpried)
			} else if result.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(TokenInvalid)
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New(TokenInvalid)
	}
	return nil, errors.New(TokenInvalid)
}

// 刷新token
func (j *JWT) RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SiginKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(7 * 24 * time.Hour).Unix()
		return j.GenerateJWT(*claims)
	}
	return "", errors.New(TokenInvalid)

}
