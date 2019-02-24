package swagger

import (
	"./Roles"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type MyCustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func getToken(email string, role Roles.Role) (string, error) {
	signingKey := []byte("EngTester")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		email,
		role.String(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
	})
	//TODO: think about roles
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("EngTester")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid {
		return token.Claims, err
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			log.Print(ve.Error())
			return token.Claims, ve
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		log.Print("Couldn't handle this token")
		return token.Claims, ve
	}
	return token.Claims, err
}
