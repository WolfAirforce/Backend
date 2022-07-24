package models

import "time"

type VipEntry struct {
	ID              int       `gorm:"column:Id;primaryKey;autoIncrement"`
	Timestamp       time.Time `gorm:"column:timestamp;default:current_timestamp"`
	PlayerName      string    `gorm:"column:playername"`
	PlayerID        string    `gorm:"column:playerid;size:20"`
	KofiEmail       string    `gorm:"column:email"`
	EndDate         time.Time `gorm:"column:enddate"`
	AdminPlayerName string    `gorm:"column:admin_playername;default:callback"`
	AdminPlayerID   string    `gorm:"column:admin_playerid;default:0:0"`
}

func (VipEntry) TableName() string {
	return "tVip"
}
