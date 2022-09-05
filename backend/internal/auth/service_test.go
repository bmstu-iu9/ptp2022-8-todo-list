package auth

import (
	"errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
	"testing"
)

type AuthTestCase struct {
	Name        string
	Input       interface{}
	Want        UserData
	IsOK        bool
	NeedToEqual bool
}

type LogoutTestCase struct {
	refreshToken string
}

type RefreshTestCase struct {
	refreshToken string
}

type LoginTestCase struct {
	Req LoginUserRequest
}

func TestAuthService(t *testing.T) {
	tokens, _ := GenerateTokens("slava@example.com", 0)
	s := service{&mockRepository{
		users: []entity.User{
			{
				Id:          0,
				Email:       "slava@example.com",
				Nickname:    "slavaruswarrior",
				Password:    "3dfff1ca8a9696f67616a2b8abd1bce3", //wasdqwertytest
				IsActivated: true,
			},
		},
		tokens: []DbToken{
			{
				userId:       0,
				refreshToken: tokens.RefreshToken,
			},
			{
				userId:       5,
				refreshToken: "token",
			},
		},
	}}
	tests := []AuthTestCase{
		{
			Name:  "login OK",
			Input: LoginTestCase{LoginUserRequest{Email: "slava@example.com", Password: "wasdqwertytest"}},
			Want: UserData{
				User:   entity.UserDto{Id: 0, Email: "slava@example.com", Nickname: "slavaruswarrior"},
				Tokens: tokens,
			},
			IsOK:        true,
			NeedToEqual: true,
		},
		{
			Name:  "login FAIL",
			Input: LoginTestCase{LoginUserRequest{Email: "sl2va@example.com", Password: "wasdqwertytest"}},
			IsOK:  false,
		},
		{
			Name:  "logout Ok",
			Input: LogoutTestCase{"token"},
			IsOK:  true,
		},
		{
			Name:  "logout Fail",
			Input: LogoutTestCase{"aaaaa"},
			IsOK:  false,
		},
		{
			Name:  "refresh Ok",
			Input: RefreshTestCase{tokens.RefreshToken},
			Want: UserData{
				User:   entity.UserDto{Id: 0, Email: "slava@example.com", Nickname: "slavaruswarrior"},
				Tokens: tokens,
			},
			IsOK:        true,
			NeedToEqual: true,
		},
		{
			Name:  "refresh FAIL",
			Input: RefreshTestCase{"non bd token"},
			Want:  UserData{},
			IsOK:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			var (
				got UserData
				err error
			)
			switch tc.Input.(type) {
			case LoginTestCase:
				got, err = s.Login(tc.Input.(LoginTestCase).Req)
			case RefreshTestCase:
				got, err = s.Refresh(tc.Input.(RefreshTestCase).refreshToken)
			case LogoutTestCase:
				err = s.Logout(tc.Input.(LogoutTestCase).refreshToken)
			}
			if tc.IsOK {
				test.IsNil(t, err)
			} else {
				test.NotNil(t, err)
			}
			if tc.NeedToEqual {
				test.DeepEqual(t, tc.Want, got)
			}
		})
	}
}

type ValidTestCase struct {
	Name    string
	IsValid bool
	token   string
	userId  int
}

func TestAccessTokenValidate(t *testing.T) {
	tokens, _ := GenerateTokens("slava@example.com", 0)
	accessTests := []ValidTestCase{
		{
			Name: "OK access", IsValid: true,
			token: tokens.AccessToken, userId: 0,
		},
		{
			Name: "FAIL access", IsValid: false,
			token: tokens.RefreshToken,
		},
	}
	for _, tc := range accessTests {
		t.Run(tc.Name, func(t *testing.T) {
			got := ValidateAccessToken(tc.token, tc.userId)
			if got != tc.IsValid {
				t.Fatalf("expected validation result: %#v, got: %#v", tc.IsValid, got)
			}
		})
	}
}

func TestRefreshTokenValidate(t *testing.T) {
	tokens, _ := GenerateTokens("slava@example.com", 0)
	refreshTests := []ValidTestCase{
		{
			Name: "OK refresh", IsValid: true,
			token: tokens.RefreshToken,
		},
		{
			Name: "FAIL refresh", IsValid: false,
			token: tokens.AccessToken,
		},
	}
	for _, tc := range refreshTests {
		t.Run(tc.Name, func(t *testing.T) {
			got := ValidateRefreshToken(tc.token)
			if got != tc.IsValid {
				t.Fatalf("expected validation result: %#v, got: %#v", tc.IsValid, got)
			}
		})
	}
}

type mockRepository struct {
	users  []entity.User
	tokens []DbToken
}

func (m mockRepository) DeleteDeadUsers() error {
	return nil
}

func (m mockRepository) GetToken(refreshToken string, userId int) (DbToken, error) {
	for _, token := range m.tokens {
		if token.refreshToken == refreshToken || token.userId == userId {
			return token, nil
		}
	}
	return DbToken{}, errors.New("no token in db")
}

func (m mockRepository) UpdateToken(userId int, newRefreshToken string) error {
	for i, token := range m.tokens {
		if token.userId == userId {
			m.tokens[i].refreshToken = newRefreshToken
			return nil
		}
	}
	return errors.New("smth wrong")
}

func (m *mockRepository) CreateToken(userId int, refreshToken string) error {
	m.tokens = append(m.tokens, DbToken{
		userId:       userId,
		refreshToken: refreshToken,
	})
	return nil
}

func (m mockRepository) DeleteToken(refreshToken string) error {
	for i, token := range m.tokens {
		if token.refreshToken == refreshToken {
			m.tokens[i] = DbToken{}
			return nil
		}
	}
	return errors.New("no token in db")
}

func (m mockRepository) GetUser(email entity.Email, userId int) (entity.User, error) {
	for _, user := range m.users {
		if int(user.Id) == userId || user.Email == email {
			return user, nil
		}
	}
	return entity.User{}, errors.New("user not found")
}
