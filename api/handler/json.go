package handler

import (
	"encoding/json"

	"firstServer/api/model"
	"fmt"
	"net/http"
)

//	handler - java의 서비스 및 dao 역활을 하는 펑션들의 모임으로 생각함.
//	차후 Redis 연결을 여기서할듯
//	연결관련된 정보는 Config 폴더를 만들어서 넣는거롤 진행 할 예정,,,

//	Header 셋팅을 공통으로 빼낼 방법은 없나? 서비스서 하는게맞나? 라우터만 처리하면되지않을까?

//	Rdb 만들기전에 redis에 먼저 데이터를 적재해볼까,,? 아직 R DB 셋팅도 안했으니,,,

// func TestJsonGet(w http.ResponseWriter) string {
func TestJsonGet(w http.ResponseWriter) {

	//json, _ := json.Marshal(model.Result{"test111", "test.com111"})

	//var r = http.ResponseWriter()
	//r.Write([]byte("리턴테스트"))

	//r, _ := redis.InitRedisClient()

	fmt.Println("#######################")
	//fmt.Println(r.Get("test1"))

	//tmp, _ := r.Get("test2").Result()
	//fmt.Println(tmp)

	fmt.Println("#######################")
	//w.Write(json)

	//return string(json)

}

func TestJsonPost(w http.ResponseWriter, r *http.Request) string {

	user := new(model.User)
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {

		fmt.Fprint(w, "Bad Request : ", err)

	}

	fmt.Println("###############################################")
	fmt.Println(user.Nickname)

	fmt.Println("###############################################")
	return ""
}

func TestJsonPatch() {

}

func TestJsonDelete() {

}

func TestMux(w http.ResponseWriter, r *http.Request) {

	fmt.Println("TEST MUX Function !!")
}
