package loginService

import "github.com/gin-gonic/gin"

type ILoginService interface {
	Login() gin.HandlerFunc
}
