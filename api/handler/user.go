package handler

import (
	"encoding/json"
	"firstServer/api/jwt"
	"firstServer/api/model"
	"firstServer/api/redis"
	"firstServer/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JoinGin(c *gin.Context) {
	var user model.User
	result := &model.Result{}

	// JSON 바인딩 (GIN의 강력한 기능)
	if err := c.ShouldBindJSON(&user); err != nil {
		result.ResultCode = "E"
		result.ResultMsg = "Parameter Exception"
		c.JSON(http.StatusBadRequest, result)
		return
	}

	rd := redis.NewRedisClient()

	if userInfo, _ := rd.HGetData("user", user.EmailId); userInfo != "" {
		result.ResultCode = "F"
		result.ResultMsg = "이미 가입되어 있습니다."
	} else {
		jsonData := util.Marshal(user)
		if err := rd.HSetData("user", user.EmailId, string(jsonData)); err != nil {
			result.ResultCode = "F"
			result.ResultMsg = "회원가입 중 오류가 발생했습니다."
		} else {
			result.ResultCode = "S"
			result.ResultMsg = "정상적으로 회원가입이 완료 되었습니다."
		}
	}

	c.JSON(http.StatusOK, result)
}

func LoginGin(c *gin.Context) {
	var user model.User
	result := &model.Result{}

	// JSON 바인딩
	if err := c.ShouldBindJSON(&user); err != nil {
		result.ResultCode = "E"
		result.ResultMsg = "Input Error"
		c.JSON(http.StatusBadRequest, result)
		return
	}

	// Redis 클라이언트
	rd := redis.NewRedisClient()
	userInfo, err := rd.HGetData("user", user.EmailId)

	if err != nil {
		result.ResultCode = "F"
		result.ResultMsg = "회원정보조회시 오류가 발생되었습니다."
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	if userInfo == "" {
		result.ResultCode = "F"
		result.ResultMsg = "회원정보가 존재하지 않습니다."
		c.JSON(http.StatusNotFound, result)
		return
	}

	// 사용자 비교
	var userDetail model.User
	if err := json.Unmarshal([]byte(userInfo), &userDetail); err != nil {
		result.ResultCode = "F"
		result.ResultMsg = "사용자 정보 파싱 오류"
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	if userDetail.Password != user.Password {
		result.ResultCode = "F"
		result.ResultMsg = "비밀번호가 일치하지 않습니다."
		c.JSON(http.StatusUnauthorized, result)
		return
	}

	// 토큰 생성
	token := jwt.MakeToken(user.EmailId)
	if token == nil {
		result.ResultCode = "F"
		result.ResultMsg = "토큰생성 오류."
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	// Redis에 리프레시 토큰 저장
	if err := rd.HSetData("userSession", user.EmailId, token.RefreshToken); err != nil {
		result.ResultCode = "F"
		result.ResultMsg = "Refresh Token Create Error"
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	token.ResultCode = "S"
	token.ResultMsg = "SUCCESS"

	c.JSON(http.StatusOK, token)
}

// package handler

// import (
// 	"encoding/json"
// 	"firstServer/api/jwt"
// 	"firstServer/api/model"
// 	"firstServer/api/redis"
// 	"firstServer/util"
// 	"fmt"
// 	"net/http"
// )

// func Join(w http.ResponseWriter, r *http.Request) {

// 	user := &model.User{}
// 	//user := new(User);
// 	result := &model.Result{}

// 	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
// 		result.ResultCode = "E"
// 		result.ResultMsg = "Parameter Exception"
// 		w.Write(util.Marshal(result))

// 		return
// 	}

// 	rd := redis.NewRedisClient()

// 	if userInfo, _ := rd.HGetData("user", user.EmailId); userInfo != "" && result.ResultCode != "E" { //	회원가입이 되어있는 경우
// 		result.ResultCode = "F"
// 		result.ResultMsg = "이미 가입되어 있습니다."

// 	} else {
// 		json := util.Marshal(user)
// 		if err := rd.HSetData("user", user.EmailId, string(json)); err != nil { //	등록된 회원이 없는 경우.
// 			result.ResultCode = "S"
// 			result.ResultMsg = "정상적으로 회원가입이 완료 되었습니다."

// 		}
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(util.Marshal(result))

// }

// func Login(w http.ResponseWriter, r *http.Request) {

// 	result := &model.Result{}

// 	fmt.Println("로그인 ")

// 	user := new(model.User)
// 	err := json.NewDecoder(r.Body).Decode(user)

// 	if err != nil { //input값 검증
// 		result.ResultCode = "E"
// 		result.ResultMsg = "Input Error"
// 		w.Write(util.Marshal(result))
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	//	Redis Client
// 	rd := redis.NewRedisClient()
// 	userInfo, err := rd.HGetData("user", user.EmailId)

// 	if err != nil {
// 		result.ResultCode = "F"
// 		result.ResultMsg = "회원정보조회시 오류가 발생되었습니다."
// 		w.Write(util.Marshal(result))
// 		return
// 	} else if userInfo == "" { //	사용자정보가 없을 경우

// 		result.ResultCode = "F"
// 		result.ResultMsg = "회원정보가 존재하지 않습니다."
// 		w.Write(util.Marshal(result))
// 		return
// 	}

// 	//	사용자 비교

// 	userDetail := model.User{}
// 	json.Unmarshal([]byte(userInfo), userDetail) //	유틸로 뺄 방법을 찾아보자.

// 	if userDetail.Password != user.Password { //	평문처리 >> 차후 base64 sha256 암호화 예정 // 비밀번호 오류
// 		result.ResultCode = "F"
// 		result.ResultMsg = "비밀번호가 일치하지 않습니다."
// 		w.Write(util.Marshal(result))
// 		return
// 	}

// 	token := jwt.MakeToken(user.EmailId)

// 	if token == nil {
// 		result.ResultCode = "F"
// 		result.ResultMsg = "토큰생성 오류."
// 		w.Write(util.Marshal(result))
// 		return
// 	}

// 	set := rd.HSetData("userSession", user.EmailId, token.RefreshToken)

// 	if set != nil {
// 		result.ResultCode = "F"
// 		result.ResultMsg = "Refresh Token Create Error"
// 		w.Write(util.Marshal(result))
// 		return
// 	}

// 	token.ResultCode = "S"
// 	token.ResultMsg = "SUCCESS"

// 	w.WriteHeader(http.StatusOK)
// 	w.Write(util.Marshal(token))

// }
