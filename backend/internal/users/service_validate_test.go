package users

import "testing"

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
				Email:    "email@test.com.",
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
				Email:           newStr("slava@example.com"),
				Nickname:        newStr("slavaruswarrior"),
				CurrentPassword: "Asjh2k123",
			},
		},
		{
			Name: "bad email", IsValid: false,
			Data: UpdateUserRequest{
				Email:           newStr("email@@test.com"),
				CurrentPassword: "Asjh2k123",
			},
		},
		{
			Name: "bad nickname", IsValid: false,
			Data: UpdateUserRequest{
				Email:           newStr("slava@example.com"),
				Nickname:        newStr("sl"),
				CurrentPassword: "Asjh2k123",
			},
		},
		{
			Name: "OK", IsValid: true,
			Data: UpdateUserRequest{
				Email:           newStr("slava@example.com"),
				Nickname:        newStr("12345"),
				CurrentPassword: "Asjh2k123",
			},
		},
		{
			Name: "bad nickname", IsValid: false,
			Data: UpdateUserRequest{
				Email:           newStr("slava@example.com"),
				Nickname:        newStr("-slavaruswarrior"),
				CurrentPassword: "Asjh2k123",
			},
		},
		{
			Name: "bad password", IsValid: false,
			Data: UpdateUserRequest{
				Email:           newStr("slava@example.com"),
				Nickname:        newStr("slavaruswarrior"),
				NewPassword:     newStr("123"),
				CurrentPassword: "DSfsfiusbn234",
			},
		},
		{
			Name: "bad password", IsValid: false,
			Data: UpdateUserRequest{
				Email:           newStr("slava@example.com"),
				Nickname:        newStr("slavaruswarrior"),
				NewPassword:     newStr("dsfkskfhs^3dsfsf"),
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
