package middlewares

import (
	"fmt"
	"net/http"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/tokenGenerator"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
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
		}

	}
}