package controllers

import (
	"net/http"

	s "airforce/internal/timer"

	"github.com/gin-gonic/gin"

	svc "airforce/cmd/api/services"
)

type GetRecordsQuery struct {
	Page int `form:"page,default=1" binding:"gte=1"`
}

func HandlerGetRecentRecords(c *gin.Context) {
	var query GetRecordsQuery

	if c.ShouldBindQuery(&query) == nil {
		rl, err := s.GetRecords(svc.Database.SurfTimer)

		if err != nil {
			c.String(http.StatusInternalServerError, "unexpected error occured")
			return
		}

		c.JSON(http.StatusOK, rl)
	} else {
		c.String(http.StatusNotAcceptable, "invalid query parameters")
	}
}
