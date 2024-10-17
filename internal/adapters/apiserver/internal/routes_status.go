package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var AppVersion = VersionRest{
	Service: "rest-gin/http",
	Version: "0.1.0",
	Build:   "1",
}

func GetVersion() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, AppVersion)
	}
}
