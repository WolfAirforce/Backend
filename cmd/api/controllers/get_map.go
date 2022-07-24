package controllers

import (
	"net/http"

	t "airforce/cmd/api/services/timer"

	"github.com/gin-gonic/gin"
)

func HandlerGetMap(c *gin.Context) {
	mapData, err := t.Timer.GetMapInformation(c.Param("name"))

	if err != nil {
		c.String(http.StatusNotAcceptable, "invalid map name was provided")
		return
	}

	c.JSON(http.StatusOK, mapData)
}
