package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	svc "airforce/cmd/api/services"

	"github.com/gin-gonic/gin"
)

func HandlerGetMapsFiles(c *gin.Context) {
	fileList, err := os.ReadDir(svc.Config.Directory.Maps)

	if err != nil {
		panic(err)
	}

	var mapFileList []string

	for _, file := range fileList {
		if !file.IsDir() {
			fileName := file.Name()
			fileExt := filepath.Ext(fileName)

			if fileExt == ".bsp" {
				mapFileList = append(mapFileList, strings.TrimSuffix(fileName, fileExt))
			}
		}
	}

	c.JSON(http.StatusOK, mapFileList)
}
