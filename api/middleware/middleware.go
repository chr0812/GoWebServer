package middleware

import (
	"firstServer/api/handler"
	"firstServer/api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS 미들웨어 - GIN 버전
func CorsMiddlewareGin() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authentication")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Preflight 요청 처리
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})
}

// 서버 공통 미들웨어 - GIN 버전
func ServerMiddlewareGin() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	})
}

// 토큰 검증 미들웨어 - GIN 버전
func TokenMiddlewareGin() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// 기존 ValidToken 로직을 GIN 컨텍스트로 수정
		flag, msg, token := handler.ValidTokenGin(c)

		if !flag {
			c.JSON(http.StatusUnauthorized, model.Result{
				ResultCode: "F",
				ResultMsg:  msg,
				ResultData: "",
			})
			c.Abort()
			return
		}

		// 새로운 토큰이 발급된 경우 (refresh token으로 access token 갱신)
		if token != nil {
			c.JSON(http.StatusOK, token)
			c.Abort()
			return
		}

		c.Next()
	})
}

// package middleware

// import (
// 	"encoding/json"
// 	"firstServer/api/handler"
// 	"firstServer/api/model"
// 	"net/http"
// )

// //	미들웨어를 기능별로 분기하여 필요한 서비스별 분리구현을 생각해본다. (ex.세션 및 토큰이 필요한 화면, 인증없이  접근할 수 있는 서비스 분리.)

// // 2025-04-28 미들웨어 데코레이터적용해봄.

// // 미들웨어 기능.
// func ServerMiddleware(next http.Handler) http.Handler {

// 	//	서버시작시 컨테이너 올라갈떄 한번 실행됨.
// 	//	DB커넥션 및 메모리 적재시에 넣으면 될 것 같음.

// 	// 들어오는 요청의 Response Header에 Content-Type을 추가
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

// 		//	인터셉터 및 각각 필터링 및 인증관련 서비스를 추가하면 될 것 같음.

// 		rw.Header().Add("Content-Type", "application/json")

// 		// 전달 받은 http.Handler를 호출한다.
// 		next.ServeHTTP(rw, r)
// 	})
// }

// // Cors 항시허용
// func CorsMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

// 		rw.Header().Set("Access-Control-Allow-Origin", "*")
// 		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		next.ServeHTTP(rw, r)
// 	})
// }

// //

// func TokenMiddleware(next http.Handler) http.Handler {

// 	//	서버시작시 컨테이너 올라갈떄 한번 실행됨.
// 	//	DB커넥션 및 메모리 적재시에 넣으면 될 것 같음.

// 	// 들어오는 요청의 Response Header에 Content-Type을 추가

// 	//	accessToken 인증 > 실패스 refreshToken 인증 가능하도록 프로세스 변경

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		flag, str, token := handler.ValidToken(w, r)

// 		if !flag {
// 			Result, _ := json.Marshal(model.Result{"F", str, ""})
// 			w.Write(Result)
// 			return
// 		} else if flag == true && token != nil {
// 			nick, _ := json.Marshal(token)
// 			w.Write(nick)
// 			return
// 		}

// 		//	인터셉터 및 각각 필터링 및 인증관련 서비스를 추가하면 될 것 같음.
// 		w.Header().Add("Content-Type", "application/json")

// 		// 전달 받은 http.Handler를 호출한다.
// 		next.ServeHTTP(w, r)
// 	})
// }
