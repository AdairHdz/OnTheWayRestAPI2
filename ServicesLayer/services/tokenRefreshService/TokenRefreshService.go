package tokenRefreshService

import (	
	"net/http"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/tokenGenerator"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

type TokenRefreshService struct{}

func (TokenRefreshService) RefreshToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := request.ParseFromRequest(context.Request, request.HeaderExtractor{"Token-Request"}, func (token *jwt.Token) (interface{}, error){
			return 	tokenGenerator.VerifyKey, nil
		}, request.WithClaims(&tokenGenerator.CustomClaim{}))
		
		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "El token no fue válido")
			return
		}

		claims, claimsConversionWentOK := token.Claims.(*tokenGenerator.CustomClaim)

		if claimsConversionWentOK && token.Valid {			
			generatedToken, err := tokenGenerator.CreateToken(claims.UserInfo.EmailAddress, claims.UserInfo.UserType)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusConflict, "Error al generar el token")
				return
			}

			response := struct {
				Token string `json:"token"`
			}{
				Token: generatedToken,	
			}

			context.JSON(http.StatusOK, response)
		} else {
			context.AbortWithStatusJSON(http.StatusConflict, "Error")
			return
		}
	}
}