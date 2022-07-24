package models

import "time"

type LatestRecords struct {
	SteamID string    `gorm:"column:steamid;primarykey" json:"steam_id"`
	Name    string    `gorm:"column:name" json:"steam_name"`
	RunTime float32   `gorm:"column:runtime" json:"run_time"`
	MapName string    `gorm:"column:map" json:"map_name"`
	Date    time.Time `gorm:"column:date;primarykey" json:"date"`
}

func (LatestRecords) TableName() string {
	return "ck_latestrecords"
}
