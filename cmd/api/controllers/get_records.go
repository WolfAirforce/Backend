package controllers

import (
	t "airforce/cmd/api/services/timer"

	"net/http"

	"github.com/gin-gonic/gin"
)

type GetRecordsQuery struct {
	Page int `form:"page,default=1" binding:"gte=1"`
}

func HandlerGetRecentRecords(c *gin.Context) {
	var query GetRecordsQuery

	if c.ShouldBindQuery(&query) == nil {
		rl, err := t.Timer.GetRecords()

		if err != nil {
			c.String(http.StatusInternalServerError, "unexpected error occured")
			return
		}

		c.JSON(http.StatusOK, rl)
	} else {
		c.String(http.StatusNotAcceptable, "invalid query parameters")
	}
}
