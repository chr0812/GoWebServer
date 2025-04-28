package main

import (
	"firstServer/api"
	"firstServer/api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

//	서버실행시 최초 메인서비스 시작, 이후 실행되어있는 라우터서버(ListenAndServe 로 메모리에 올려둔 서버? 가 실행되어있다~)
//	라우터로 url 호출시엔 메인펑션은 실행되지 않으며, 라우터로만 진입되는걸 볼 수 있다
//

//	미들웨어 설정
//	CROS는 모든 도메인 허용, 차후 등록도메인 리스트 처리해야할까?
//	 - router.use() 를 사용하여 미들웨어를 등록함.( 전체 요청에 대한 등록으로 예상 )
//	JWT 토큰적용, 차후 리플래쉬 토큰 및 인증관련해서 추가예정
//	 - 토큰이 필요한 URL요청에 대해 직접적으로 적용.
//	Server middleware 생성, request, response 에 공통으로 등록되어야할 가능 추가하면 됨.
//	 - 현재까진 디테일하게 등록하지 않음.

func main() {
	/*
		//1.  서버 온
		api.RegisterRouters()

		//	2. DB연결유무 체크, 서버실행시 서버 풀 연결처리?
		fmt.Println("스타트?")
		port := ":9000"

		server := &http.Server{ //(2) server 객체 할당
			Addr:              port,
			ReadHeaderTimeout: time.Duration(60 * time.Second),
			WriteTimeout:      time.Duration(60 * time.Second),
			Handler:           http.DefaultServeMux,
		}

		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("라우터서버연결실패")
		}
	*/

	router := mux.NewRouter()
	api.MuxRouters(router)

	//router.Use(mux.CORSMethodMiddleware(router))
	router.Use(middleware.CorsMiddleware)   //	CROS 허용 미들웨어	-
	router.Use(middleware.ServerMiddleware) //	Server설정 미들웨어

	http.ListenAndServe(":9000", router)

}
