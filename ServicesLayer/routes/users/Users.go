package users

import (
	"github.com/gin-gonic/gin"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/userManagementService"
)

var (
	_userManagementService = userManagementService.UserManagementService{}
)

func Routes(route *gin.RouterGroup) {
	users := route.Group("/users")
	{
		users.PATCH("/:userId/verify", _userManagementService.VerifyAccount())
		users.POST("/:userId/verify", _userManagementService.GetNewVerificationCode())
	}
}