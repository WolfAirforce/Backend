package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"airforce/cmd/api/services/vip"
	"airforce/internal/vip/response"
)

type FormData struct {
	Data response.KofiCallbackData `form:"data" binding:"required"`
}

func PostKofiCallback(c *gin.Context) {
	var f FormData

	if err := c.ShouldBind(&f); err != nil {
		c.String(http.StatusBadRequest, "invalid data")
		return
	}

	if os.Getenv("KOFI_VERIFICATION_TOKEN") != f.Data.VerificationToken {
		c.String(http.StatusBadRequest, "invalid verification token was given")
		return
	}

	u, err := vip.Manager.UpdateUserFromData(f.Data)

	if err != nil {
		vip.Manager.Webhook.CreateContent(fmt.Sprintf("Encountered an error with a donation: `%s` (%s)", err.Error(), f.Data.KofiTransactionID))
		c.String(http.StatusOK, "err")
	} else {
		vip.Manager.Webhook.CreateContent(fmt.Sprintf("Received a new donation from `%s` for $%.2f. Updated VIP end date to %s.", u.PlayerID, f.Data.Amount, u.EndDate))
		c.String(http.StatusOK, "ok")
	}
}
