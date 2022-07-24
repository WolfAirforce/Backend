package models

type MapTier struct {
	MapName     string `gorm:"column:mapname;primarykey" json:"map_name"`
	Tier        int    `gorm:"column:tier" json:"tier"`
	MaxVelocity int    `gorm:"column:maxvelocity" json:"max_velocity"`
}

func (MapTier) TableName() string {
	return "ck_maptier"
}
