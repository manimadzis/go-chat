package jwtmanager

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

type JWTAuth struct {
	JWT                   string    `json:"jwt"`
	RefreshToken          string    `json:"refresh_tokenToken"`
	JWTExpiredAt          time.Time `json:"jwt_expired_at"`
	RefreshTokenExpiredAt time.Time `json:"refresh_token_expired_at"`
}

type Config struct {
	SignKey            string
	RefreshTokenLength int
	JWTTTL             time.Duration
	RefreshTokenTTL    time.Duration
}

type JWTManager interface {
	NewRefreshToken() (string, time.Time)
	NewJWTAuth(userId string) (*JWTAuth, error)
	RefreshJWT(userId string, refreshToken string) (*JWTAuth, error)
	Parse(token string) (string, error)
}
type jwtManager struct {
	signKey            string
	refreshTokenLength int
	JWTTTL             time.Duration
	refreshTokenTTL    time.Duration
}

func (j jwtManager) RefreshJWT(userId string, refreshToken string) (*JWTAuth, error) {
	var err error
	auth := JWTAuth{
		RefreshToken: refreshToken,
	}
	auth.JWT, err = j.newJWT(userId, j.JWTTTL)
	if err != nil {
		return nil, err
	}
	auth.JWTExpiredAt = time.Now().Add(j.JWTTTL)
	return &auth, nil
}

func (j jwtManager) newJWT(userId string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   userId,
		ExpiresAt: time.Now().Add(duration).Unix(),
	})

	return token.SignedString([]byte(j.signKey))
}

func (j jwtManager) newRefreshToken() string {
	token := make([]byte, j.refreshTokenLength)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Read(token)
	return base64.StdEncoding.EncodeToString(token)
}

func (j jwtManager) parse(token string) (string, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(j.signKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := jwtToken.Claims.(jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("failed to get claims from token")
	}

	return claims.Subject, nil
}

func (j jwtManager) NewJWTAuth(userId string) (*JWTAuth, error) {
	var err error
	jwtToken := JWTAuth{}
	jwtToken.JWT, err = j.newJWT(userId, j.JWTTTL)
	if err != nil {
		return nil, err
	}
	jwtToken.RefreshToken = j.newRefreshToken()
	jwtToken.RefreshTokenExpiredAt = time.Now().Add(j.refreshTokenTTL)
	return &jwtToken, nil
}

func (j jwtManager) NewRefreshToken() (string, time.Time) {
	return j.newRefreshToken(), time.Now().Add(j.refreshTokenTTL)
}

func (j jwtManager) Parse(token string) (string, error) {
	return j.parse(token)
}

func New(config Config) (JWTManager, error) {
	if config.SignKey == "" {
		return nil, ErrEmptySignKey
	}

	if config.RefreshTokenTTL == time.Duration(0) {
		config.RefreshTokenTTL = time.Duration(30 * 24 * time.Hour)
	}
	if config.JWTTTL == time.Duration(0) {
		config.RefreshTokenTTL = time.Duration(15 * time.Minute)
	}
	if config.RefreshTokenLength == 0 {
		config.RefreshTokenLength = 32
	}

	return &jwtManager{
		signKey:            config.SignKey,
		refreshTokenLength: config.RefreshTokenLength,
		JWTTTL:             config.JWTTTL,
		refreshTokenTTL:    config.RefreshTokenTTL,
	}, nil
}
