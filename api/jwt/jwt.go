package jwt

import (
	"firstServer/api/model"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateRefreshToken(Name string) string {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["Name"] = Name
	claims["exp"] = time.Now().Add(time.Hour * 30).Unix()
	t, err := accessToken.SignedString([]byte("refresh"))

	if err != nil {
		return ""
	}
	return t

}

func CreateAccessToken(Name string) string {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["Name"] = Name
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	t, err := accessToken.SignedString([]byte("access"))

	if err != nil {
		return ""
	}
	return t
}

/*
	func ExtractToken(r string) string {
		fmt.Println(r)
		strArr := strings.Split(r, " ")
		fmt.Println(len(strArr))

		if len(strArr) == 2 {

			return strArr[1]

		}
		return ""
	}
*/

func VerifyToken(r string, s string) error {

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(r, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s), nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

// 토큰 생성
func MakeToken(emailId string) *model.Token {

	token := new(model.Token)
	token.AccessToken = CreateAccessToken(emailId)

	if token.AccessToken == "" {
		fmt.Println("Create Access Token Error")
		return nil
	}

	token.RefreshToken = CreateRefreshToken(emailId)

	if token.RefreshToken == "" {
		fmt.Println("Create Refresh Token Error")
		return nil
	}
	return token
}
