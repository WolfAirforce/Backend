package controllers

import (
	"net/http"

	s "airforce/internal/timer"

	"github.com/gin-gonic/gin"

	svc "airforce/cmd/api/services"
)

func HandlerGetMaps(c *gin.Context) {
	mapList, err := s.GetAllMapInformation(svc.Database.SurfTimer)

	if err != nil {
		c.String(http.StatusInternalServerError, "unexpected error occured")
		return
	}

	c.JSON(http.StatusOK, mapList)
}
