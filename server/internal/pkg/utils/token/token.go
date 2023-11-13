package token

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func MakeToken(userId int) (string, error) {
	tokenExpiredAt, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRED_AT"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenExpiredAt)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
}

func ExtractToken(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func IsValidToken(ctx *gin.Context) error {
	tokenString := ExtractToken(ctx)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})
	if err != nil {
		log.Printf("err: %v", err)
		return err
	}
	return nil
}

func ExtractTokenID(ctx *gin.Context) (int, error) {
	tokenString := ExtractToken(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		return 0, fmt.Errorf("failed to parse")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id, err := strconv.ParseInt(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return int(id), nil
	}
	return 0, nil
}
