package routes

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/middlewares"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/routes/providers"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/routes/requesters"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/routes/serviceRequests"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/routes/states"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/loginService"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/registerService"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/tokenRefreshService"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/logoutService"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	_loginService = loginService.LoginService{}
	_registerService = registerService.RegisterService{}
	_tokenRefreshService = tokenRefreshService.TokenRefreshService{}
	_logoutService = logoutService.LogoutService{}
)


func setupLogOutput() {
	f, _ := os.Create("./ServicesLayer/logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func init(){
	setupLogOutput()
	router = gin.Default()
	router.Use(middlewares.Logger())
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	limiter := tollbooth.NewLimiter(50, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	limiter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	v1 := router.Group("/v1", tollbooth_gin.LimitHandler(limiter))	
	{
				
		router.StaticFS("/images", http.Dir("./images"))
		router.StaticFS("/reviews", http.Dir("./public/reviews"))
		v1.POST("/register", _registerService.RegisterUser())
		v1.POST("/login", _loginService.Login())		
		requesters.Routes(v1)
		providers.Routes(v1)
		states.Routes(v1)
		serviceRequests.Routes(v1)
	}
	
	refresh := v1.Group("/refresh")
	{
		refresh.Use(middlewares.AuthenticateWithRefreshToken())
		refresh.POST("/", _tokenRefreshService.RefreshToken())
	}

	logout := v1.Group("/logout")
	{
		logout.Use(middlewares.Authenticate())
		logout.POST("/", _logoutService.Logout())
	}
}

func StartServer(){
	fmt.Println("Server listening on port 8080")
	router.Run("0.0.0.0:8080")
	
	
}