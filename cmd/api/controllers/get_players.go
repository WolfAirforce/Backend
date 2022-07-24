package controllers

import (
	t "airforce/cmd/api/services/timer"

	"net/http"

	"github.com/gin-gonic/gin"
)

type GetPlayersQuery struct {
	SortBy    string `form:"sortBy,default=points"`
	SortOrder string `form:"sortOrder,default=desc"`
	Page      int    `form:"page,default=1" binding:"gte=1"`
}

func HandlerGetPlayers(c *gin.Context) {
	var query GetPlayersQuery

	if c.ShouldBindQuery(&query) == nil {
		pl, err := t.Timer.GetPlayers(query.SortBy, query.SortOrder, 25, (query.Page-1)*25)

		if err != nil {
			c.String(http.StatusInternalServerError, "unexpected error occured")
			return
		}

		c.JSON(http.StatusOK, pl)
	} else {
		c.String(http.StatusNotAcceptable, "invalid query parameters")
	}
}
