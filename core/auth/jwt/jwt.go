package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var (
	key = []byte("123456")
	issuer = "test"
)

type ApiClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

type Auth struct {
	Token string `json:"token"`
	ExpireIn string  `json:"expire_at"`
}

func Init()  {
	key = []byte(viper.GetString("jwt.key"))
	issuer = viper.GetString("jwt.issuer")
}

func GetToken(userId int) (Auth, error) {
	claims := &ApiClaims{
		userId,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),
			ExpiresAt: int64(time.Now().Unix() + (60 * 60 * 24)),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		return Auth{}, err
	}
	return Auth{Token:ss, ExpireIn: strconv.Itoa(60 * 60 * 24)}, nil
}

func CheckToken(tokenStr string) (userId int, isValid bool)  {
	token, err := jwt.ParseWithClaims(tokenStr, &ApiClaims{}, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, false
	}
	if claims, ok := token.Claims.(*ApiClaims); ok && token.Valid {
		return claims.UserId, true
	} else {
		return 0, false
	}
}