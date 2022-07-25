package main

import (
	"airforce/cmd/api/controllers"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	svc "airforce/cmd/api/services"
)

func main() {
	p, e := os.LookupEnv("CFG_FILE_PATH")

	if !e {
		p = "/tmp/config.json"
	}

	svc.Initialize(p)

	gin.SetMode(svc.Config.Server.Mode)

	r := gin.Default()
	r.SetTrustedProxies(svc.Config.Server.TrustedProxies)

	s := persistence.NewInMemoryStore(time.Minute)

	r.GET("/maps", cache.CachePage(s, time.Minute*5, controllers.HandlerGetMaps))
	r.GET("/maps/:name", cache.CachePage(s, time.Minute*5, controllers.HandlerGetMap))
	r.GET("/maps/files", cache.CachePage(s, time.Minute*5, controllers.HandlerGetMapsFiles))

	r.GET("/players", controllers.HandlerGetPlayers)
	r.GET("/players/:id", controllers.HandlerGetPlayerByID)

	r.GET("/records/recent", controllers.HandlerGetRecentRecords)
	// /api/records/maps/<map>
	// /api/records/players/<player>
	// /api/records/players/<player>/<map>

	r.POST("/callback/kofi", controllers.PostKofiCallback)

	r.Run(fmt.Sprintf(":%d", svc.Config.Server.Port))
}
