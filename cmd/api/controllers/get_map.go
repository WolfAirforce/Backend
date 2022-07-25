package controllers

import (
	"net/http"

	svc "airforce/cmd/api/services"
	s "airforce/internal/timer"

	"github.com/gin-gonic/gin"
)

func HandlerGetMap(c *gin.Context) {
	mapData, err := s.GetMapInformation(svc.Database.SurfTimer, c.Param("name"))

	if err != nil {
		c.String(http.StatusNotAcceptable, "invalid map name was provided")
		return
	}

	c.JSON(http.StatusOK, mapData)
}
