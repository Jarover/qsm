//routes/routes.go

package routes

import (
	"github.com/Jarover/qsm/cmd/qsm/pages"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/", pages.StartPage)
	return r
}
