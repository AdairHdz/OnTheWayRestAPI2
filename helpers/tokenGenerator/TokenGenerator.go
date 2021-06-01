package tokenGenerator

import (
	"crypto/rsa"
	"io/ioutil"
	"time"
	"github.com/dgrijalva/jwt-go"		
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
		EmailAddress string
		UserType int
	}
}

func CreateToken(emailAddress string, userType int) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	token.Claims = &CustomClaim{
		&jwt.StandardClaims{			
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
		struct{EmailAddress string; UserType int}{
			EmailAddress: emailAddress,
			UserType: userType,
		},
	}

	return token.SignedString(signKey)
}

func CreateRefreshToken(emailAddress string, userType int) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	token.Claims = &CustomClaim{
		&jwt.StandardClaims{			
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
		struct{EmailAddress string; UserType int}{
			EmailAddress: emailAddress,
			UserType: userType,
		},
	}

	return token.SignedString(signKey)
}


