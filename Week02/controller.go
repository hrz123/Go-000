package Week02

import (
	"code.byted.org/gopkg/logs"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hrz123/Go-000/Week02/code"
	"net/http"
	"strconv"
)

type Response struct {
	Code    int32
	Message string
	Data    interface{}
}

func GetUserInfo(c *gin.Context) {
	fmt.Println(c.Param("userID"))
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		logs.Error("param user id cannot be parsed")
		c.JSON(http.StatusOK, &Response{
			Code:    1,
			Message: "user id 参数错误",
			Data:    nil,
		})
		return
	}
	userInfo, err := DealWithUser(int64(userID))
	if err != nil {
		if errors.Is(err, code.NotFound) {
			// dao层的错误
			c.JSON(http.StatusOK, &Response{
				Code:    1,
				Message: "dao error",
				Data:    nil,
			})
			return
		}
	}
	c.JSON(200, &Response{
		Code:    0,
		Message: "success",
		Data:    userInfo,
	})
}
