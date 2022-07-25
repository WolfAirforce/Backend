package vip

import (
	"airforce/internal/vip/database/models"
	"airforce/internal/vip/response"
	"errors"
	"math"
	"regexp"
	"time"

	"gorm.io/gorm"
)

var (
	steamIdRegex = regexp.MustCompile(`STEAM_[0-5]:([01]:\d+)`)
)

func TimeFromAmount(amount float64) time.Duration {
	return time.Duration(24*31*int(math.Floor(amount/3.0))) * time.Hour
}

func UpdateUserFromData(db *gorm.DB, data response.KofiCallbackData) (models.VipEntry, error) {
	v := models.VipEntry{}
	d := TimeFromAmount(data.Amount)

	if err := db.Where("email = ?", data.Email).First(&v).Error; err != nil {
		// try to find their steam id in their msg
		reMatch := steamIdRegex.FindStringSubmatch(data.Message)

		if reMatch != nil {
			v = models.VipEntry{
				PlayerName: "Kofi Customer",
				PlayerID:   reMatch[1],
				KofiEmail:  data.Email,
				EndDate:    time.Now().Add(d),
			}

			db.Create(&v)
		} else {
			// we will need to find them and ask for their steam id, which will then be added manually.
			return v, errors.New("steam id was not provided")
		}
	} else {
		// bought before so just update
		var newTime time.Time

		if v.EndDate.Before(time.Now()) {
			newTime = time.Now().Add(d)
		} else {
			newTime = v.EndDate.Add(d)
		}

		db.Model(&models.VipEntry{}).Where("email = ?", data.Email).Update("enddate", newTime)
		db.Where("email = ?", data.Email).First(&v)
	}

	db.Commit()

	return v, nil
}
