package api

import (
	"firstServer/api/handler"
	"firstServer/api/middleware"

	"github.com/gin-gonic/gin"
)

func GinRouters(r *gin.Engine) {
	// 인증이 필요 없는 API
	public := r.Group("/")
	{
		public.GET("/api/json", handler.TestGin)
		public.POST("/api/createToken", handler.CreateTokenGin)
		public.POST("/api/authenticate", handler.CreateTokenGin)
	}

	// 사용자 관련 API
	user := r.Group("/user")
	{
		user.POST("/join", handler.JoinGin)
		user.POST("/login", handler.LoginGin)
	}

	// 토큰 인증이 필요한 API
	auth := r.Group("/")
	auth.Use(middleware.TokenMiddlewareGin())
	{
		auth.GET("/api/token", handler.TestGin)
		auth.GET("/chart/aChart", handler.GetAChartGin)
	}
}

// package api

// import (
// 	"firstServer/api/handler"
// 	"firstServer/api/middleware"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// // 라우터로 각 서비스별 url 호출에 대한 대응 및 Handler 로 리턴역활을 한다.
// // java spring의 Controller 역활만 수행하도록 함
// // 1차는 여러가지 기능 테스트를 위해 function 을 직접구현 후 각각 분기처리 하도록 한다.

// //

// func MuxRouters(router *mux.Router) {

// 	//	서버 접근테스트

// 	//	JSON
// 	router.Handle("/api/json", http.HandlerFunc(handler.TestMux))

// 	//	Token 검증관련
// 	router.Handle("/api/token", middleware.TokenMiddleware(http.HandlerFunc(handler.TestMux)))

// 	router.Handle("/chart/aChart", middleware.TokenMiddleware(http.HandlerFunc(handler.GetAChart)))

// 	//	인증이 필요없는화면
// 	router.Handle("/api/createToken", http.HandlerFunc(handler.CreateToken)) //.Methods("POST")

// 	router.Handle("/api/authenticate", http.HandlerFunc(handler.CreateToken)) //.Methods("POST")

// 	router.Handle("/user/join", http.HandlerFunc(handler.Join))
// 	router.Handle("/user/login", middleware.CorsMiddleware(http.HandlerFunc(handler.Login)))

// }

// /*
// func RegisterRouters() {

// 	fmt.Println("라우터 진입.")

// 	//	서버 접근테스트

// 	//json테스트
// 	http.HandleFunc("/api/json", func(w http.ResponseWriter, rq *http.Request) {

// 		user := &model.User{}

// 		switch rq.Method {
// 		case http.MethodGet:
// 			//	get 방식 호출로 요청시 함께 전달된 Parameter은 Url 에서 추출한다.
// 			fmt.Println("Get 호출 테스트 완료")
// 			fmt.Println(rq.URL.Query().Get("nickname"))
// 			nick, _ := json.Marshal(user)
// 			w.Write(nick)
// 		case http.MethodPost:
// 			fmt.Println("Post 호출 테스트 완료")
// 			err := json.NewDecoder(rq.Body).Decode(&user)

// 			if err != nil {
// 				w.WriteHeader(http.StatusBadRequest)
// 				fmt.Println("잘못된요청")
// 				return
// 			}
// 		case http.MethodPatch:
// 			fmt.Println("Patch 호출 테스트 완료")
// 			err := json.NewDecoder(rq.Body).Decode(&user)

// 			if err != nil {
// 				w.WriteHeader(http.StatusBadRequest)
// 				fmt.Println("잘못된요청")
// 				return
// 			}
// 		case http.MethodDelete:
// 			fmt.Println("Delete 호출 테스트 완료")
// 			fmt.Println(rq.URL.Query().Get("nickname"))
// 		}
// 	})

// 	//json테스트
// 	http.HandleFunc("/api/token", func(w http.ResponseWriter, rq *http.Request) {

// 		//user := &model.User{}

// 		switch rq.Method {
// 		case http.MethodGet:

// 		case http.Meth																																																																																																																														odPost:
// 			handler.CreateToken(w, rq)
// 		case http.MethodPatch:

// 		case http.MethodDelete:
// 		}

// 	})

// 	//json테스트
// 	http.HandleFunc("/api/tokenValidation", func(w http.ResponseWriter, rq *http.Request) {

// 		//user := &model.User{}

// 		switch rq.Method {
// 		case http.MethodGet:

// 		case http.MethodPost:
// 			handler.ValidToken(w, rq)
// 		case http.MethodPatch:

// 		case http.MethodDelete:
// 		}

// 	})

// }
// */
// //	badRequest 리턴
