package models

type PlayerRank struct {
	CountryName string `gorm:"column:country" json:"country"`
	Steam       struct {
		ID   string `gorm:"column:steamid;primarykey" json:"id"`
		ID64 string `gorm:"column:steamid64" json:"id64"`
		Name string `gorm:"column:name" json:"name"`
	} `gorm:"embedded" json:"steam"`
	Points struct {
		Overall                uint32 `gorm:"column:points" json:"overall"`
		WorldRecordPoints      uint32 `gorm:"column:wrpoints" json:"world_record"`
		BonusWorldRecordPoints uint32 `gorm:"column:wrbpoints" json:"bonus_world_record"`
		StageWorldRecordPoints uint32 `gorm:"column:wrcppoints" json:"stage_world_record"`
		MapTopTenPoints        uint32 `gorm:"column:top10points" json:"map_top_ten"`
		GroupPoints            uint32 `gorm:"column:groupspoints" json:"group"`
		MapPoints              uint32 `gorm:"column:mappoints" json:"map"`
		BonusPoints            uint32 `gorm:"column:bonuspoints" json:"bonus"`
	} `gorm:"embedded" json:"points"`
	Completion struct {
		FinishedMaps    uint32 `gorm:"column:finishedmaps" json:"maps"`
		FinishedMapsPro uint32 `gorm:"column:finishedmapspro" json:"maps_pro"`
		FinishedBonuses uint32 `gorm:"column:finishedbonuses" json:"bonuses"`
		FinishedStages  uint32 `gorm:"column:finishedstages" json:"stages"`
	} `gorm:"embedded" json:"completion"`
	Records struct {
		WorldRecords      uint32 `gorm:"column:wrs" json:"world"`
		BonusWorldRecords uint32 `gorm:"column:wrbs" json:"bonus_world"`
		StageWorldRecords uint32 `gorm:"column:wrcps" json:"stage_world"`
		MapTopTens        uint32 `gorm:"column:top10s" json:"top_ten"`
	} `gorm:"embedded" json:"records"`
	Activity struct {
		TimeSpent struct {
			TimeSpentAlive      uint32 `gorm:"column:timealive" json:"alive"`
			TimeSpentSpectating uint32 `gorm:"column:timespec" json:"spectating"`
		} `gorm:"embedded" json:"time_spent"`
		LastSeen        uint64 `gorm:"column:lastseen" json:"last_seen_date"`
		Joined          uint64 `gorm:"column:joined" json:"join_date"`
		ConnectionCount uint32 `gorm:"column:connections" json:"connections"`
	} `gorm:"embedded" json:"activity"`
	Groups uint32 `gorm:"column:groups" json:"groups"`
	Style  uint16 `gorm:"column:style" json:"style"`
}

func (PlayerRank) TableName() string {
	return "ck_playerrank"
}
