package middlewares

import (
	"fmt"
	"net/http"

	tokenBlackList "github.com/AdairHdz/OnTheWayRestAPI/helpers/tokenBlackList"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/tokenGenerator"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func Authenticate() gin.HandlerFunc {
	return func(context *gin.Context){
		token, err := request.ParseFromRequest(context.Request, request.OAuth2Extractor, func (token *jwt.Token) (interface{}, error){			
			return 	tokenGenerator.VerifyKey, nil
		}, request.WithClaims(&tokenGenerator.CustomClaim{}))		

		if err != nil {
			fmt.Println("Invalid token", err)			
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		extractedToken, err := request.OAuth2Extractor.ExtractToken(context.Request)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusForbidden, "Error while trying to extract token")
			return
		}

		tokenBlackListHandler := tokenBlackList.GetInstance()
		_, err = tokenBlackListHandler.Get(fmt.Sprintf("BlackListedToken_%v", extractedToken))
		if err != nil {
			if err == redis.Nil {
				context.Next()
				return
			}
			context.AbortWithStatusJSON(http.StatusForbidden, "There was an error while trying to validate your token")
			return
		}		
		context.AbortWithStatusJSON(http.StatusForbidden, "This token can no longer be used")
		return		
	}
}

func AuthenticateWithRefreshToken()  gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := request.ParseFromRequest(context.Request, request.HeaderExtractor{"Token-Request"}, func (token *jwt.Token) (interface{}, error){
			return 	tokenGenerator.VerifyKey, nil
		}, request.WithClaims(&tokenGenerator.CustomClaim{}))
						

		if err != nil {
			fmt.Println("Invalid token", err)
			context.AbortWithStatus(http.StatusUnauthorized)			
			return
		}

		if !token.Valid {
			context.AbortWithStatus(http.StatusUnauthorized)			
			return
		}	
		
		extractedRefreshToken, err := request.HeaderExtractor{"Token-Request"}.ExtractToken(context.Request)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusForbidden, "Error while trying to extract token")
			return
		}

		tokenBlackListHandler := tokenBlackList.GetInstance()
		_, err = tokenBlackListHandler.Get(fmt.Sprintf("BlackListedRefreshToken_%v", extractedRefreshToken))
		if err != nil {
			if err == redis.Nil {
				context.Next()
				return
			}
			context.AbortWithStatusJSON(http.StatusForbidden, "There was an error while trying to validate your token")
			return
		}		
		context.AbortWithStatusJSON(http.StatusForbidden, "This token can no longer be used")
		return	
	}
}