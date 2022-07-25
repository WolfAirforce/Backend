package services

import (
	"log"

	"airforce/internal/config"

	"gorm.io/gorm"
)

var (
	Database struct {
		SurfTimer *gorm.DB
		VIP       *gorm.DB
	}
	Config *config.Config
)

func Initialize(filePath string) error {
	log.Printf("Initializing services using configuration in %s\n", filePath)

	// initialize cfg
	var err error
	Config, err = config.NewConfig(filePath)

	if err != nil {
		panic(err)
	}

	// initialize databases
	Database.SurfTimer, err = CreateConnection(Config.Database.SurfTimer)

	if err != nil {
		return err
	}

	Database.VIP, err = CreateConnection(Config.Database.VIP)

	return err
}
