package token

import (
	"errors"
	"server/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

type TokenClaims struct {
	UserID   uint      `json:"user_id"`
	Username string    `json:"username"`
	Type     TokenType `json:"type"`
	jwt.RegisteredClaims
}

// GenerateTokenPair creates both access and refresh tokens for the given user
func GenerateTokenPair(user model.UserModel, accessSecret, refreshSecret string) (accessToken, refreshToken string, err error) {
	// Generate access token
	accessToken, err = generateToken(user, TokenTypeAccess, time.Hour*24, accessSecret) // 24 hours
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshToken, err = generateToken(user, TokenTypeRefresh, time.Hour*24*7, refreshSecret) // 7 days
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// 生成token
func generateToken(user model.UserModel, tokenType TokenType, expiration time.Duration, jwtSecret string) (string, error) {
	// Create the Claims
	claims := TokenClaims{
		UserID:   user.ID,
		Username: user.Username,
		Type:     tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the token and returns the claims
func ValidateToken(tokenString, jwtSecret string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// RefreshToken generates a new access token using a refresh token
func RefreshToken(refreshTokenString, jwtSecret string) (string, error) {
	// Validate the refresh token
	claims, err := ValidateToken(refreshTokenString, jwtSecret)
	if err != nil {
		return "", err
	}

	// Verify that this is actually a refresh token
	if claims.Type != TokenTypeRefresh {
		return "", errors.New("invalid token type")
	}

	// Create a new user object with the claims data
	user := model.UserModel{
		Username: claims.Username,
	}
	user.ID = claims.UserID

	// Generate a new access token
	return generateToken(user, TokenTypeAccess, time.Hour*24, jwtSecret)
}
