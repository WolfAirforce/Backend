package services

import (
	"airforce/internal/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateConnection(cfg config.ConfigDatabase) (*gorm.DB, error) {
	uri := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	return gorm.Open(mysql.Open(uri), &gorm.Config{})
}
