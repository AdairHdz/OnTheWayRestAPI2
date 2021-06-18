package tokenRefreshService

import "github.com/gin-gonic/gin"

type ITokenRefreshService interface {
	RefreshToken() gin.HandlerFunc
}
