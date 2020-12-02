package Week02

import (
	"code.byted.org/gopkg/logs"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.Form["user_id"][0])
	if err != nil {
		logs.Error("form value user id cannot be parsed")
	}
	userInfo, err := DealWithUser(int64(userID))
	if err != nil {
		logs.Error("get user info from DB error %+v", err)
	}
	resp, err := json.Marshal(userInfo)
	if err != nil {
		logs.Error("json marshal error %+v", err)
	}
	_, _ = w.Write(resp)
}
