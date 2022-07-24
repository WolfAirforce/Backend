package timer

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"airforce/internal/steam"
	tdb "airforce/internal/timer/database"
	"airforce/internal/timer/database/models"
)

type SurfTimer struct {
	Database *gorm.DB
}

func (st SurfTimer) Connect(ci *tdb.ConnectionInformation) error {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", ci.Username, ci.Password, ci.Host, ci.Port, ci.DatabaseName)
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})

	if err != nil {
		return err
	}

	st.Database = db

	return nil
}

/* Player functions */
func (st SurfTimer) GetPlayerByID(steamId string) (player models.PlayerRank, err error) {
	if !steam.IsSteamID(steamId) {
		return models.PlayerRank{}, errors.New("provided steam id is not valid")
	}

	res := st.Database.First(&player, "steamid = ?", steamId)
	err = res.Error

	return
}

func (st SurfTimer) GetPlayerByID64(steamId64 string) (player models.PlayerRank, err error) {
	if !steam.IsSteamID64(steamId64) {
		return models.PlayerRank{}, errors.New("provided steam id is not valid")
	}

	res := st.Database.First(&player, "steamid64 = ?", steamId64)
	err = res.Error

	return
}

func (st SurfTimer) GetPlayers(sortBy string, sortOrder string, limit int, offset int) (playerList []models.PlayerRank, err error) {
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

	res := st.Database.Order(fmt.Sprintf("%s %s", sortBy, sortOrder)).Limit(limit).Offset(offset).Find(&playerList)
	err = res.Error

	return
}

/* Records functions */
func (st SurfTimer) GetRecords() (rl []models.LatestRecords, err error) {
	res := st.Database.Limit(10).Offset(0).Order("date desc").Find(&rl)
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

func (st SurfTimer) GetAllMapInformation() (ml []models.CustomMap, err error) {
	res := st.Database.Raw(baseMapQuery).Scan(&ml)
	err = res.Error

	return
}

func (st SurfTimer) GetMapInformation(mapName string) (m models.CustomMap, err error) {
	res := st.Database.Raw(singleMapQuery, mapName).Scan(&m)
	err = res.Error

	return
}
