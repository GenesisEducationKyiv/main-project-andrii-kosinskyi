package model_test

import (
	"testing"

	"bitcoin_checker_api/internal/model"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  string
	}{
		{name: "Valid creating user", email: "taras@shcevchenko.com", want: "taras@shcevchenko.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tUser := model.NewUser(tt.email); tUser.Email != tt.want {
				t.Errorf("TestNewUser() name = %v got = %s want = %s", tt.name, tUser.Email, tt.want)
			}
		})
	}
}
