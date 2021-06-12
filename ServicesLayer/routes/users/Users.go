package users

import (
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/userManagementService"
	"github.com/gin-gonic/gin"
)

var (
	_userManagementService = userManagementService.UserManagementService{}
)

func Routes(route *gin.RouterGroup) {
	users := route.Group("/users")
	{
		users.PATCH("/:userId/verify", _userManagementService.VerifyAccount())
		users.PUT("/:userId/verify", _userManagementService.GetNewVerificationCode())
		users.PATCH("/password", _userManagementService.RecoverPassword())
		users.PUT("/recoveryCode", _userManagementService.SendRecoveryCode())
	}
}
