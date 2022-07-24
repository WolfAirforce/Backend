package controllers

import (
	t "airforce/cmd/api/services/timer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerGetPlayerByID(c *gin.Context) {
	p, err := t.Timer.GetPlayerByID64(c.Param("id"))

	if err != nil {
		c.String(http.StatusNotAcceptable, "invalid steam id was provided")
		return
	}

	c.JSON(http.StatusOK, p)
}
