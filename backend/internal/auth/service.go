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
	Login(email, password string) (entity.UserDto, Token, error)
	Logout(refreshToken string) error
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

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

func (s service) Login(email, password string) (entity.UserDto, Token, error) {
	entityUser, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return entity.UserDto{}, Token{}, errors.New("user not found")
	}
	isPassEquals := entityUser.Password == fmt.Sprintf("%x", md5.Sum([]byte(password)))
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

func (s service) SaveRefreshToken(userId int, refreshToken string) error {
	ok, err := s.repo.CheckToken(userId)
	if err != nil {
		return err
	}
	if ok {
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

func GenerateTokens(email string) (Token, error) {
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
