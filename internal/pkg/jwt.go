package pkg

import (
	"CSR/internal/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
)

type CustomClaims struct {
	jwt.StandardClaims
	UserID    int `json:"user_id"`
	Role models.Role `json:"role"`
	IsRefresh bool `json:"isRefresh"`

}

func GenerateToken(userId, ttl int, isRefresh bool,role models.Role) (string,error) {
	logger:=zerolog.New(os.Stdout).With().Timestamp().Logger()
	if isRefresh{
		logger.Info().Msg("generating refresh token")
	}else{
		logger.Info().Msg("generating access token")
	}
	claims := CustomClaims{StandardClaims: jwt.StandardClaims{},
		UserID:    userId,
		IsRefresh: isRefresh,
		Role:role,
	}
	log.Println("isRefresh",isRefresh)
	if isRefresh {
		claims.StandardClaims.ExpiresAt = int64(time.Duration(ttl) * 24 * time.Hour)
		log.Println("claims.StandardClaims.ExpiresAt (refresh token)",claims.StandardClaims.ExpiresAt)
	} else {
		claims.StandardClaims.ExpiresAt = int64(time.Duration(ttl) *5* time.Minute)
		log.Println("claims.StandardClaims.ExpiresAt (access token)",claims.StandardClaims.ExpiresAt)
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString,_:=token.SignedString([]byte(os.Getenv("JWT_KEY")))
	return tokenString,nil
}

func ParseToken(tokenString string) (int,bool,models.Role,error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method:%v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return 0,false,"",err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		log.Println("token String ,claims.UserID,claims.isRefresh",tokenString,claims.UserID,claims.IsRefresh)
		return claims.UserID,claims.IsRefresh, claims.Role,nil
	}
	return 0, false,"",fmt.Errorf("invalid token")
}
