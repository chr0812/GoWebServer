package handler

import (
	"encoding/json"
	"fmt"

	"firstServer/api/model"
	"firstServer/api/redis"
	"net/http"

	"firstServer/api/jwt"
)

/*

	핸들러로의 토큰 함수와, util 형태의 토큰 함수 를 분기하는 것에 대한 고민이 필요
		-	미들웨어로 사용중인 핸들러 함수는 단독으로만 실행가능.
		-	타 영역에서의 접근이 가능하게 하기 위한 util 형태의 함수도 필요할 수 있다고 생각함.
*/

// 토큰 생성요청.
func CreateToken(w http.ResponseWriter, r *http.Request) {

	user := new(model.User)

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.Write([]byte("error"))
		return
	}

	token := jwt.MakeToken(user.EmailId)

	fmt.Println(token)

	if token == nil {
		w.Write([]byte("CreateToken Error"))
		return
	}
	//	Redis
	//	nickName 기준으로 refreshToken 값을 메모리에 저장.

	rd := redis.NewRedisClient()

	result := rd.HSetData("userSession", user.EmailId, token.RefreshToken)

	if result != nil {
		w.Write([]byte("Refresh Token Setting Error"))
		return
	}

	nick, _ := json.Marshal(token)

	w.Write(nick)

}

func ValidToken(w http.ResponseWriter, r *http.Request) (bool, string, *model.Token) {

	token := r.Header.Get("authentication") //	헤더 토큰값 받기.

	//	1. 엑세스토큰확인.
	access := jwt.VerifyToken(token, "access")

	if access != nil { //	엑세스 토큰이 정상이 아닌경우.

		mToken := new(model.Token)
		err := json.NewDecoder(r.Body).Decode(mToken)

		if err != nil { //	입력값 파싱 오류
			return false, "Parameter Decode Error", nil
		}

		rd := redis.NewRedisClient()
		str, err := rd.HGetData("userSession", mToken.EmailId)

		if err != nil {

			if str == token { //	refresh token 검증완료 >> access token 발급.
				mToken := jwt.MakeToken(mToken.EmailId)
				return true, "", mToken
			} else { //	refresh token 검증실패 >> 재 로그인안내
				return false, "", nil
			}
		}
	}

	return true, "", nil
}

// Access 토큰 확인.
func ValidAccessToken(w http.ResponseWriter, r *http.Request) (bool, string) {
	//	1. 엑세스토큰 확인 >>
	//	2. 엑세스토큰 오류시 오류 리턴

	mToken := new(model.Token)
	err := json.NewDecoder(r.Body).Decode(mToken)

	if err != nil { //	json 형태의 입력값을 model(VO) 에 바인딩.
		return false, "Parameter Decode Error"
	}

	err2 := jwt.VerifyToken(mToken.AccessToken, "access") //	차후 config or property 형태로 변경

	if err2 != nil {
		return false, "Access Token Validate"
	}

	return true, "Success"

}

//	Refrash 토큰 확인.

func ValidRefreshToken(w http.ResponseWriter, r *http.Request) (bool, string) {
	mToken := new(model.Token)
	err := json.NewDecoder(r.Body).Decode(mToken)

	if err != nil {
		return false, "Parameter Decode Error"

	}

	err3 := jwt.VerifyToken(mToken.RefreshToken, "refresh") //	차후 config or property 형태로 변경

	if err3 != nil {

		//
		return false, "refresh Token Validate"
	}

	return true, "Success"
}
