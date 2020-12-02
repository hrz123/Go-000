package Week02

import (
	"database/sql"
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

func getUserFromDB(userID int64) (*UserInfo, error) {
	userInfo := new(UserInfo)
	db := h.Slave()
	result := db.Table("user_info").Where("id = ?", userID).Scan(&userInfo)
	if result.Error != nil {
		if result.Error == sql.ErrNoRows {
			return nil, errors.Wrapf(result.Error, "No Row found for user id %d", userID)
		}
		return nil, errors.Wrap(result.Error, "Other Dao error")
	}
	return userInfo, nil
}
