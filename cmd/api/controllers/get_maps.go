package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	t "airforce/cmd/api/services/timer"
)

func HandlerGetMaps(c *gin.Context) {
	mapList, err := t.Timer.GetAllMapInformation()

	if err != nil {
		c.String(http.StatusInternalServerError, "unexpected error occured")
		return
	}

	c.JSON(http.StatusOK, mapList)
}
