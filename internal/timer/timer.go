package timer

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"airforce/internal/steam"
	"airforce/internal/timer/database/models"
)

/* Player functions */
func GetPlayerByID(db *gorm.DB, steamId string) (player models.PlayerRank, err error) {
	if !steam.IsSteamID(steamId) {
		return models.PlayerRank{}, errors.New("provided steam id is not valid")
	}

	res := db.First(&player, "steamid = ?", steamId)
	err = res.Error

	return
}

func GetPlayerByID64(db *gorm.DB, steamId64 string) (player models.PlayerRank, err error) {
	if !steam.IsSteamID64(steamId64) {
		return models.PlayerRank{}, errors.New("provided steam id is not valid")
	}

	res := db.First(&player, "steamid64 = ?", steamId64)
	err = res.Error

	return
}

func GetPlayers(db *gorm.DB, sortBy string, sortOrder string, limit int, offset int) (playerList []models.PlayerRank, err error) {
	sortBy = strings.ToLower(sortBy)
	sortOrder = strings.ToLower(sortOrder)

	switch sortBy {
	case "points", "timealive", "wrs", "wrbs", "wrcps":
		break
	default:
		return nil, errors.New("sort key must be one of points, timealive, wrs, wrbs, or wrcps")
	}

	switch sortOrder {
	case "asc", "desc":
		break
	default:
		return nil, errors.New("sort order must be one of asc, desc")
	}

	res := db.Order(fmt.Sprintf("%s %s", sortBy, sortOrder)).Limit(limit).Offset(offset).Find(&playerList)
	err = res.Error

	return
}

/* Records functions */
func GetRecords(db *gorm.DB) (rl []models.LatestRecords, err error) {
	res := db.Limit(10).Offset(0).Order("date desc").Find(&rl)
	err = res.Error

	return
}

/* Map functions */
const (
	baseMapQuery = `
		SELECT
		mapname,
		tier,
		maxvelocity,
		(
			SELECT
				COUNT(*) + 1 
			FROM
				ck_zones 
			WHERE
				ck_zones.mapname = ck_maptier.mapname 
				AND ck_zones.zonegroup = 0 
				AND ck_zones.zonetype = 3 
		)
		AS stages,
		(
			SELECT
				COUNT(*) 
			FROM
				(
					SELECT DISTINCT
						mapname,
						zonegroup 
					FROM
						ck_zones 
					WHERE
						ck_zones.mapname = ck_maptier.mapname 
						AND zonegroup > 0 
				)
				AS BZ 
		)
		AS bonuses,
		(
			SELECT
				CASE
					WHEN
						EXISTS
						(
							SELECT
								* 
							FROM
								ck_zones 
							WHERE
								ck_zones.mapname = ck_maptier.mapname 
						)
					THEN
						1 
					ELSE
						0 
				END
		)
		AS is_zoned 
	FROM
		ck_maptier
	`
	singleMapQuery = baseMapQuery + " WHERE ck_maptier.mapname = ?"
)

func GetAllMapInformation(db *gorm.DB) (ml []models.CustomMap, err error) {
	res := db.Raw(baseMapQuery).Scan(&ml)
	err = res.Error

	return
}

func GetMapInformation(db *gorm.DB, mapName string) (m models.CustomMap, err error) {
	res := db.Raw(singleMapQuery, mapName).Scan(&m)
	err = res.Error

	return
}
