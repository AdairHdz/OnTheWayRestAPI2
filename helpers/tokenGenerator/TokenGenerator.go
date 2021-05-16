package tokenGenerator

import (
	"crypto/rsa"
	"io/ioutil"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"	
)


const (
	privKeyPath = "./helpers/tokenGenerator/app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "./helpers/tokenGenerator/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	VerifyKey  *rsa.PublicKey
	signKey    *rsa.PrivateKey
)

func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)

	if err != nil {
		panic("There was an error while trying to read the private key")
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)

	if err != nil {
		panic("There was an error while trying to parse the private key")
	}
	
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)

	if err != nil {
		panic("There was an error while trying to read the public key")
	}

	
	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)

	if err != nil {
		panic("There was an error while trying to parse the public key")
	}

}

type CustomClaim struct {
	*jwt.StandardClaims
	UserInfo struct{
		Username string
		UserType int
	}
}

func CreateToken(username string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	token.Claims = &CustomClaim{
		&jwt.StandardClaims{			
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
		struct{Username string; UserType int}{
			Username: username,
			UserType: 1,
		},
	}

	return token.SignedString(signKey)
}

func ValidateToken(){
	request.ParseFromRequest(nil, request.OAuth2Extractor, func (token *jwt.Token) (interface{}, error){
		return 	VerifyKey, nil
	}, request.WithClaims(&CustomClaim{}))
}