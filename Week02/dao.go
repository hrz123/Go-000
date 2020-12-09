package Week02

import (
	"fmt"
	"github.com/hrz123/Go-000/Week02/code"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type HDB struct {
	*gorm.DB            // Master
	SlaveDB  []*gorm.DB // Slaves
}

func (h *HDB) Master() *gorm.DB {
	return h.DB
}

func (h *HDB) Slave() *gorm.DB {
	if len(h.SlaveDB) == 0 {
		return h.DB
	}
	return h.SlaveDB[0]
}

var h = new(HDB)

type UserInfo struct {
	Id     int64  `gorm:"primary_key"`
	Name   string `gorm:"size:50"`
	Gender bool
}

func GetUserFromDB(userID int64) (*UserInfo, error) {
	userInfo := new(UserInfo)
	db := h.Slave()
	result := db.Table("user_info").Where("id = ?", userID).Scan(&userInfo)
	if result.Error != nil {
		return nil, errors.Wrapf(code.NotFound, "sql: %s error: %v",
			fmt.Sprintf("select * from user_info where id = %d", userID), result.Error)
	}
	return userInfo, nil
}
