package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var mapFileList []string

func init() {
	m, exists := os.LookupEnv("MAP_FOLDER")

	if !exists {
		m = "/maps"
	}

	files, err := os.ReadDir(m)

	if err != nil {
		panic(err)
	}

	for _, item := range files {
		if !item.IsDir() {
			fileName := item.Name()
			fileExt := filepath.Ext(fileName)

			if fileExt == ".bsp" {
				mapFileList = append(mapFileList, strings.TrimSuffix(fileName, fileExt))
			}
		}
	}
}

func HandlerGetMapsFiles(c *gin.Context) {
	c.JSON(http.StatusOK, mapFileList)
}
