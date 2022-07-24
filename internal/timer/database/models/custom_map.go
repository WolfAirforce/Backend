package models

type CustomMap struct {
	MapName     string `gorm:"column:mapname;primarykey" json:"name"`
	Tier        int    `gorm:"column:tier" json:"tier"`
	MaxVelocity int    `gorm:"column:maxvelocity" json:"max_velocity"`
	Stages      int    `gorm:"column:stages" json:"stages"`
	Bonuses     int    `gorm:"column:bonuses" json:"bonuses"`
	IsZoned     bool   `gorm:"column:is_zoned" json:"is_zoned"`
}
