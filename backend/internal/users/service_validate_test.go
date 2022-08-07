package users

import (
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
)

func TestCreateUserRequest_Validate(t *testing.T) {
	tests := []struct {
		Name    string
		IsValid bool
		Data    CreateUserRequest
	}{
		{
			Name: "OK", IsValid: true,
			Data: CreateUserRequest{
				Email:    "slava@example.com",
				Nickname: "slavaruswarrior",
				Password: "Asjh2k123",
			}},
		{
			Name: "bad email", IsValid: false,
			Data: CreateUserRequest{
				Email:    "email@@test.com",
				Nickname: "slavaruswarrior",
				Password: "Asjh2k123",
			}},
		{
			Name: "bad nickname", IsValid: false,
			Data: CreateUserRequest{
				Email:    "slava@example.com",
				Nickname: "sl",
				Password: "Asjh2k123",
			}},
		{
			Name: "OK", IsValid: true,
			Data: CreateUserRequest{
				Email:    "slava@example.com",
				Nickname: "12345",
				Password: "Asjh2k123",
			}},
		{
			Name: "bad nickname", IsValid: false,
			Data: CreateUserRequest{
				Email:    "slava@example.com",
				Nickname: "-slavaruswarrior",
				Password: "Asjh2k123",
			}},
		{
			Name: "bad password", IsValid: false,
			Data: CreateUserRequest{
				Email:    "slava@example.com",
				Nickname: "slavaruswarrior",
				Password: "123",
			}},
		{
			Name: "bad password", IsValid: false,
			Data: CreateUserRequest{
				Email:    "slava@example.com",
				Nickname: "slavaruswarrior",
				Password: "dsfkskfhs^3dsfsf",
			}},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			got := tc.Data.Validate()

			if (got == nil) != tc.IsValid {
				t.Fatalf("expected validation result: %#v, got: %#v", tc.IsValid, got)
			}
		})
	}
}

func TestUpdateUserRequest_Validate(t *testing.T) {
	tests := []struct {
		Name    string
		IsValid bool
		Data    UpdateUserRequest
	}{
		{
			Name: "OK", IsValid: true,
			Data: UpdateUserRequest{
				Email:           entity.NewEmail("slava@example.com"),
				Nickname:        entity.NewNickname("slavaruswarrior"),
				CurrentPassword: "Asjh2k123",
			},
		},
		{
			Name: "bad email", IsValid: false,
			Data: UpdateUserRequest{
				Email:           entity.NewEmail("email@@test.com"),
				CurrentPassword: "Asjh2k123",
			},
		},
		{
			Name: "bad nickname", IsValid: false,
			Data: UpdateUserRequest{
				Email:           entity.NewEmail("slava@example.com"),
				Nickname:        entity.NewNickname("sl"),
				CurrentPassword: "Asjh2k123",
			},
		},
		{
			Name: "OK", IsValid: true,
			Data: UpdateUserRequest{
				Email:           entity.NewEmail("slava@example.com"),
				Nickname:        entity.NewNickname("12345"),
				CurrentPassword: "Asjh2k123",
			},
		},
		{
			Name: "bad nickname", IsValid: false,
			Data: UpdateUserRequest{
				Email:           entity.NewEmail("slava@example.com"),
				Nickname:        entity.NewNickname("-slavaruswarrior"),
				CurrentPassword: "Asjh2k123",
			},
		},
		{
			Name: "bad password", IsValid: false,
			Data: UpdateUserRequest{
				Email:           entity.NewEmail("slava@example.com"),
				Nickname:        entity.NewNickname("slavaruswarrior"),
				NewPassword:     entity.NewPassword("123"),
				CurrentPassword: "DSfsfiusbn234",
			},
		},
		{
			Name: "bad password", IsValid: false,
			Data: UpdateUserRequest{
				Email:           entity.NewEmail("slava@example.com"),
				Nickname:        entity.NewNickname("slavaruswarrior"),
				NewPassword:     entity.NewPassword("dsfkskfhs^3dsfsf"),
				CurrentPassword: "DSfsfiusbn234",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			got := tc.Data.Validate()

			if (got == nil) != tc.IsValid {
				t.Fatalf("expected validation result: %#v, got: %#v", tc.IsValid, got)
			}
		})
	}
}
