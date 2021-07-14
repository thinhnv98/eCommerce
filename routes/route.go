package routes

import (
	"database/sql"

	"eCommerce/controllers"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Db     *sql.DB
	Server *gin.Engine
}

func (_self Route) Register() {
	//Controller
	common := controllers.Common{}
	commonRoutes := _self.Server.Group("/api")
	{
		commonRoutes.GET("/ping", common.Ping)
	}
}