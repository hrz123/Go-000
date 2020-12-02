package Week02

import (
	"code.byted.org/gopkg/logs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Response struct {
	Code    int32
	Message string
	Data    interface{}
}

func GetUserInfo(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		logs.Error("param user id cannot be parsed")
		c.JSON(http.StatusOK, &Response{
			Code:    0,
			Message: "user id 参数错误",
			Data:    nil,
		})
		return
	}
	userInfo, err := DealWithUser(int64(userID))
	if err != nil {
		logs.Error("get user info from DB error %+v", err)
		c.JSON(http.StatusOK, &Response{
			Code:    0,
			Message: "get user info from db error",
			Data:    nil,
		})
	}
	c.JSON(200, &Response{
		Code:    0,
		Message: "success",
		Data:    userInfo,
	})
}
