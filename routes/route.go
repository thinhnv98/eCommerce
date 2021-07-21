package routes

import (
	"database/sql"
	"eCommerce/repository"
	"eCommerce/service"

	"eCommerce/controller"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Db     *sql.DB
	Server *gin.Engine
}

func (_self Route) Register() {
	//Controller
	common := controller.Common{}
	commonRoutes := _self.Server.Group("/api/ping")
	{
		commonRoutes.GET("", common.Ping)
	}

	user := controller.UserController{
		IUserService: service.UserService{
			IUserRepo: repository.UserRepo{
				Db: _self.Db,
			},
		},
	}
	userRoutes := _self.Server.Group("/api/user")
	{
		userRoutes.POST("/create", user.Create)
		userRoutes.POST("/login", user.Login)
	}
}
