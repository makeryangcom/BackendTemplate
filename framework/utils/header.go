package Utils

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func CheckHeader(c *gin.Context) (bool, int) {

	status := false
	accountUID := 0

	referer := c.Request.Header.Get("Client-Referer")
	if referer != "" {
		allowedReferer := []string{
			"",
		}
		for _, allowedReferer := range allowedReferer {
			if strings.HasPrefix(referer, allowedReferer) {
				status = true
			}
		}
	}

	token := c.Request.Header.Get("Client-Token")
	if token != "" {
		uidMap, _ := DecodeId(128, token)
		if len(uidMap) > 0 {
			accountUID = uidMap[0]
		}
	}

	return status, accountUID
}
