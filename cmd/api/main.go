package main

import (
	"airforce/cmd/api/controllers"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.0.186"})

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

	r.Run(":8080")
}
