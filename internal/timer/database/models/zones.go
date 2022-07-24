package models

type Zones struct {
	MapName string `gorm:"column:mapname;primarykey" json:"map_name"`
}

func (Zones) TableName() string {
	return "ck_zones"
}
