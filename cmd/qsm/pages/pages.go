package pages

import (
	"net/http"

	"github.com/Jarover/qsm/cmd/qsm/config"
	"github.com/gin-gonic/gin"
)

// StartPage
func StartPage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"version": config.Version.VersionStr(),
		"data":    config.Version.BuildTime,
	})
}
