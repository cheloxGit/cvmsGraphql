package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cheloxGit/cvmsGraphql/resolvers"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

var jwtSecret []byte = []byte("thepolyglotdeveloper")

//ValidateJWT func
func ValidateJWT(t string) (interface{}, error) {
	if t == "" {
		return nil, errors.New("Authorization token must be present")
	}
	token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var decodedToken interface{}
		mapstructure.Decode(claims, &decodedToken)
		return decodedToken, nil
	} else {
		return nil, errors.New("Invalid authorization token")
	}
}

//CreateTokenEndpoint func
func CreateTokenEndpoint(response http.ResponseWriter, request *http.Request) {
	var user resolvers.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})
	tokenString, error := token.SignedString(jwtSecret)
	if error != nil {
		fmt.Println(error)
	}
	response.Header().Set("content-type", "application/json")
	response.Write([]byte(`{ "token": "` + tokenString + `" }`))
}
