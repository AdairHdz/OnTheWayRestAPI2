package logoutService

import "github.com/gin-gonic/gin"

type ILogoutService interface {
	Logout() gin.HandlerFunc
}
