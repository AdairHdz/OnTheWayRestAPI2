package registerService

import "github.com/gin-gonic/gin"

type IRegisterService interface {
	RegisterUser() gin.HandlerFunc
}
