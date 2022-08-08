package auth

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

const JWT_ACCESS_SECRET = "secret-key"
const JWT_REFRESH_SECRET = "refresh-key"

type Service interface {
	SaveRefreshToken(userId int, refreshToken string) error
	Login(email entity.Email, password entity.Password) (entity.UserDto, Token, error)
	Logout(refreshToken string) error
	Refresh(refreshToken string) (entity.UserDto, Token, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type DbToken struct {
	userId       int    `json:"user_id"`
	refreshToken string `json:"refresh_token"`
}

type LoginUserRequest struct {
	Email    entity.Email    `json:"email"`
	Password entity.Password `json:"password"`
}

type Claims struct {
	Email entity.Email
	jwt.RegisteredClaims
}

func (s service) Login(email entity.Email, password entity.Password) (entity.UserDto, Token, error) {
	entityUser, err := s.repo.GetUser(email, -1)
	if err != nil {
		return entity.UserDto{}, Token{}, errors.New("user not found")
	}
	isPassEquals := entityUser.Password == *entity.NewPassword(fmt.Sprintf("%x", md5.Sum([]byte(password))))
	if !isPassEquals {
		return entity.UserDto{}, Token{}, errors.New("incorrect password")
	}
	user := entity.NewUserDto(entityUser)
	tokens, err := GenerateTokens(user.Email)
	if err != nil {
		return entity.UserDto{}, Token{}, err
	}
	err = s.SaveRefreshToken(int(user.Id), tokens.RefreshToken)
	return user, tokens, err
}

func (s service) Logout(refreshToken string) error {
	return s.repo.DeleteToken(refreshToken)
}

func (s service) Refresh(refreshToken string) (entity.UserDto, Token, error) {
	if refreshToken == "" {
		return entity.UserDto{}, Token{}, errors.New("no token")
	}
	isTokenValid := validateRefreshToken(refreshToken)
	tokenFromDb, err := s.repo.GetToken(refreshToken)
	if !isTokenValid || err != nil {
		return entity.UserDto{}, Token{}, errors.New("wrong token")
	}
	entityUser, err := s.repo.GetUser("", tokenFromDb.userId)
	if err != nil {
		return entity.UserDto{}, Token{}, err
	}
	user := entity.NewUserDto(entityUser)
	tokens, err := GenerateTokens(user.Email)
	if err != nil {
		return entity.UserDto{}, Token{}, err
	}
	err = s.SaveRefreshToken(int(user.Id), tokens.RefreshToken)
	return user, tokens, err
}

func (s service) SaveRefreshToken(userId int, refreshToken string) error {
	_, err := s.repo.GetToken(refreshToken)
	if err == nil {
		err = s.repo.UpdateToken(userId, refreshToken)
		if err != nil {
			return err
		}
	} else {
		err = s.repo.CreateToken(userId, refreshToken)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateTokens(email entity.Email) (Token, error) {
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{time.Now().Add(5 * time.Minute)},
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString([]byte(JWT_ACCESS_SECRET))
	if err != nil {
		return Token{}, err
	}
	claims.ExpiresAt = &jwt.NumericDate{time.Now().Add(30 * 24 * time.Hour)}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString([]byte(JWT_REFRESH_SECRET))
	if err != nil {
		return Token{}, err
	}
	return Token{AccessToken: accessTokenString, RefreshToken: refreshTokenString}, nil
}

func validateAccessToken(accessToken string) bool {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(JWT_ACCESS_SECRET), nil
	})
	if err != nil {
		return false
	}
	return token != nil
}

func validateRefreshToken(refreshToken string) bool {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(JWT_REFRESH_SECRET), nil
	})
	if err != nil {
		return false
	}
	return token != nil
}
