package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var mapFilePath string

func init() {
	var e bool

	mapFilePath, e = os.LookupEnv("MAP_FOLDER")

	if !e {
		mapFilePath = "/maps"
	}

	if _, err := os.Stat(mapFilePath); os.IsNotExist(err) {
		panic(err)
	}
}

func HandlerGetMapsFiles(c *gin.Context) {
	files, err := os.ReadDir(mapFilePath)

	if err != nil {
		panic(err)
	}

	var mapFileList []string

	for _, item := range files {
		if !item.IsDir() {
			fileName := item.Name()
			fileExt := filepath.Ext(fileName)

			if fileExt == ".bsp" {
				mapFileList = append(mapFileList, strings.TrimSuffix(fileName, fileExt))
			}
		}
	}

	c.JSON(http.StatusOK, mapFileList)
}
