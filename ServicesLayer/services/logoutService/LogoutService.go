package logoutService

import (
	"fmt"
	"net/http"
	"time"

	tokenBlackList "github.com/AdairHdz/OnTheWayRestAPI/helpers/tokenBlackList"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

type LogoutService struct{}

var (
	tokenBlackListHandler = tokenBlackList.GetInstance()
)

func (LogoutService) Logout() gin.HandlerFunc {
	return func(context *gin.Context) {
		extractedToken, err := request.OAuth2Extractor.ExtractToken(context.Request)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid token")
			return
		}

		err = tokenBlackListHandler.Save(fmt.Sprintf("BlackListedToken_%v", extractedToken), extractedToken, time.Minute*15)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "Error while trying to log user out")
			return
		}

		extractedRefreshToken, err := request.HeaderExtractor{"Token-Request"}.ExtractToken(context.Request)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid token")
			return
		}

		err = tokenBlackListHandler.Save(fmt.Sprintf("BlackListedRefreshToken_%v", extractedRefreshToken), extractedRefreshToken, time.Hour*168)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "Error while trying to log user out")
			return
		}

	}
}
