package swagger

import jwt "github.com/dgrijalva/jwt-go"

func getToken(email, role string) (string, error) {
	signingKey := []byte("EngTester")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
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
	return token.Claims, err
}
