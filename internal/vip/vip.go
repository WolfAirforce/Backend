package vip

import (
	"airforce/internal/vip/database/models"
	"airforce/internal/vip/response"
	"errors"
	"math"
	"regexp"
	"time"

	"github.com/disgoorg/disgo/webhook"
	"gorm.io/gorm"
)

var (
	steamIdRegex = regexp.MustCompile(`STEAM_[0-5]:([01]:\d+)`)
)

type Manager struct {
	Database *gorm.DB
	Webhook  webhook.Client
}

func TimeFromAmount(amount float64) time.Duration {
	return time.Duration(24*31*int(math.Floor(amount/3.0))) * time.Hour
}

func (m Manager) UpdateUserFromData(data response.KofiCallbackData) (models.VipEntry, error) {
	v := models.VipEntry{}
	d := TimeFromAmount(data.Amount)

	if err := m.Database.Where("email = ?", data.Email).First(&v).Error; err != nil {
		// try to find their steam id in their msg
		reMatch := steamIdRegex.FindStringSubmatch(data.Message)

		if reMatch != nil {
			v = models.VipEntry{
				PlayerName: "Kofi Customer",
				PlayerID:   reMatch[1],
				KofiEmail:  data.Email,
				EndDate:    time.Now().Add(d),
			}

			m.Database.Create(&v)
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

		m.Database.Model(&models.VipEntry{}).Where("email = ?", data.Email).Update("enddate", newTime)
		m.Database.Where("email = ?", data.Email).First(&v)
	}

	m.Database.Commit()

	return v, nil
}
