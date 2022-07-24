package timer

import (
	surfTimer "airforce/internal/timer"
	"fmt"

	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Timer surfTimer.SurfTimer

func init() {
	res, err := strconv.ParseInt(os.Getenv("DB_SURF_PORT"), 10, 16)

	if err != nil {
		panic(err)
	}

	uri := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		os.Getenv("DB_SURF_USERNAME"),
		os.Getenv("DB_SURF_PASSWORD"),
		os.Getenv("DB_SURF_HOST"),
		uint16(res),
		os.Getenv("DB_SURF_NAME"),
	)

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Timer = surfTimer.SurfTimer{
		Database: db,
	}

	if err != nil {
		panic(err)
	}
}
