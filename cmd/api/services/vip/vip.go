package vip

import (
	"airforce/internal/vip"
	"airforce/internal/vip/database/models"
	"fmt"
	"os"
	"strconv"

	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/snowflake/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Manager vip.Manager

func init() {
	res, err := strconv.ParseInt(os.Getenv("DB_VIP_PORT"), 10, 16)

	if err != nil {
		panic(err)
	}

	uri := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		os.Getenv("DB_VIP_USERNAME"),
		os.Getenv("DB_VIP_PASSWORD"),
		os.Getenv("DB_VIP_HOST"),
		uint16(res),
		os.Getenv("DB_VIP_NAME"),
	)

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.VipEntry{})

	if !db.Migrator().HasColumn(&models.VipEntry{}, "email") {
		db.Migrator().AddColumn(&models.VipEntry{}, "email")
	}

	res, err = strconv.ParseInt(os.Getenv("KOFI_WEBHOOK_ID"), 10, 64)

	if err != nil {
		panic(err)
	}

	wh := webhook.New(snowflake.ID(res), os.Getenv("KOFI_WEBHOOK_SECRET"))

	Manager = vip.Manager{
		Database: db,
		Webhook:  wh,
	}
}
