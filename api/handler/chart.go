package handler

import (
	"firstServer/api/model"
	"firstServer/api/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAChartGin(c *gin.Context) {
	result := &model.Result{}
	rd := redis.NewRedisClient()

	st, err := rd.HGetData("statistics", "aChart")
	if err != nil {
		result.ResultCode = "F"
		result.ResultMsg = "데이터 조회 오류"
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result.ResultCode = "S"
	result.ResultMsg = "데이터 조회 성공"
	result.ResultData = st

	c.JSON(http.StatusOK, result)
}

// package handler

// import (
// 	"firstServer/api/model"
// 	"firstServer/api/redis"
// 	"firstServer/util"
// 	"fmt"
// 	"net/http"
// )

// func GetAChart(w http.ResponseWriter, r *http.Request) {

// 	result := &model.Result{}
// 	rd := redis.NewRedisClient()

// 	st, err := rd.HGetData("statistics", "aChart")

// 	fmt.Println(st)

// 	if err != nil {
// 		result.ResultCode = "F"
// 		result.ResultMsg = "데이터 조회 오류"
// 		w.Write(util.Marshal(result))
// 		return
// 	}

// 	result.ResultCode = "S"
// 	result.ResultMsg = "데이터 조회 성공"
// 	result.ResultData = st

// 	w.Write(util.Marshal(result))

// }
