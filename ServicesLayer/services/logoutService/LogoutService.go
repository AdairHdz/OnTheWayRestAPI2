package logoutService

import (
	"fmt"
	"net/http"
	"time"
	tokenBlackList "github.com/AdairHdz/OnTheWayRestAPI/helpers/tokenBlackList"	
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"	
)

type LogoutService struct{ }

var (
	tokenBlackListHandler = tokenBlackList.GetInstance()
)

func (LogoutService) Logout() gin.HandlerFunc {
	return func(context *gin.Context) {
		extractedToken, err := request.OAuth2Extractor.ExtractToken(context.Request)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "Error while trying to extract token")
			return
		}
						
		err = tokenBlackListHandler.Save(fmt.Sprintf("BlackListedToken_%v", extractedToken), extractedToken, time.Minute * 15)		
		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "Errorr while trying to log user out")
		}

	}
}