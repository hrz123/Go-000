package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func main() {
	r := gin.New()
	r.GET("/:userID/", func(c *gin.Context) {
		fmt.Printf("get user info from DB error %+v", errors.Wrapf(sql.ErrNoRows,
			"No Row found for user id %d: %s", 3, sql.ErrNoRows))
	})
	_ = r.Run(":9000")
}
