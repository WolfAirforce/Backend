package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"airforce/internal/vip"
	"airforce/internal/vip/response"

	"github.com/gin-gonic/gin"

	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/snowflake/v2"

	svc "airforce/cmd/api/services"
)

func createWebhook(url string) webhook.Client {
	r := regexp.MustCompile(`\/(\d+)\/([^?&]+)`)
	m := r.FindStringSubmatch(url)

	if m == nil {
		return nil
	}

	res, _ := strconv.Atoi(m[1])

	return webhook.New(
		snowflake.ID(res),
		m[2],
	)
}

func PostKofiCallback(c *gin.Context) {
	var f struct {
		Data response.KofiCallbackData `form:"data" binding:"required"`
	}

	if err := c.ShouldBind(&f); err != nil {
		c.String(http.StatusBadRequest, "invalid data")
		return
	}

	if svc.Config.Secrets.Kofi.VerificationToken != f.Data.VerificationToken {
		c.String(http.StatusBadRequest, "invalid verification token was given")
		return
	}

	w := createWebhook(svc.Config.Notifications.Kofi.OnPayment)
	u, err := vip.UpdateUserFromData(svc.Database.VIP, f.Data)

	if err != nil {
		if w != nil {
			w.CreateContent(fmt.Sprintf("Encountered an error with a donation: `%s` (%s)", err.Error(), f.Data.KofiTransactionID))
		}

		c.String(http.StatusOK, "err")
	} else {
		if w != nil {
			w.CreateContent(fmt.Sprintf("Received a new donation from `%s` for $%.2f. Updated VIP end date to %s.", u.PlayerID, f.Data.Amount, u.EndDate))
		}

		c.String(http.StatusOK, "ok")
	}
}
