package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mahdibouaziz/go-testing-playground/webapp/pkg/data"
)

const jwtTokenExpiry = time.Minute * 15
const refreshTokenExpiry = time.Hour * 24

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	UserName string `json:"name"`
	jwt.RegisteredClaims
}

func (app *application) getTokenFromHeaderAndVerify(w http.ResponseWriter, r *http.Request) (string, *Claims, error) {
	// add a header
	w.Header().Add("Vary", "Authorization")

	// get the authorization header
	authHeader := r.Header.Get("Authorization")

	// sanity check
	if authHeader == "" {
		return "", nil, errors.New("no auth header")
	}

	// split the header on spaces
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return "", nil, errors.New("invalid auth header")
	}

	// check to see if we have the word "Bearer"
	if headerParts[0] != "Bearer" {
		return "", nil, errors.New("unauthorized: no Bearer")
	}

	// get the token
	tokenString := headerParts[1]

	// declare an empty Clais variable
	claims := &Claims{}

	// parse the token with our claims
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		hmacSampleSecret := []byte(app.JWTSecret)
		return hmacSampleSecret, nil
	})
	// The error catches expired token as well
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return "", nil, errors.New("expired token")
		}
		return "", nil, err
	}

	// make sure that we isseud this token
	if claims.Issuer != app.Domain {
		return "", nil, errors.New("incorrect issuer")
	}

	// valid token
	return tokenString, claims, nil
}

func (app *application) generateTokenPair(user *data.User) (TokenPairs, error) {
	// create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		"sub":  fmt.Sprint(user.ID),
		"aud":  app.Domain,
		"iss":  app.Domain,
		"exp":  time.Now().Add(jwtTokenExpiry).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(app.JWTSecret))
	if err != nil {
		return TokenPairs{}, err
	}

	// create the refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": fmt.Sprint(user.ID),
		"exp": time.Now().Add(refreshTokenExpiry).Unix(),
	})

	// sign the refresh token
	refreshTokenString, err := refreshToken.SignedString([]byte(app.JWTSecret))
	if err != nil {
		return TokenPairs{}, err
	}

	tokenPairs := TokenPairs{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
	}
	return tokenPairs, nil
}
