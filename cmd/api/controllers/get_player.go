package controllers

import (
	s "airforce/internal/timer"
	"net/http"

	"github.com/gin-gonic/gin"

	svc "airforce/cmd/api/services"
)

func HandlerGetPlayerByID(c *gin.Context) {
	p, err := s.GetPlayerByID64(svc.Database.SurfTimer, c.Param("id"))

	if err != nil {
		c.String(http.StatusNotAcceptable, "invalid steam id was provided")
		return
	}

	c.JSON(http.StatusOK, p)
}
