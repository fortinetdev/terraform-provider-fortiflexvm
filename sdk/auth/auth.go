package auth

import (
	"fmt"
	"os"
)

// Auth describes the authentication information for FlexVM
type Auth struct {
	Username string
	Password string
	Token    string
}

// NewAuth inits Auth object with the given metadata
func NewAuth(username, password string) *Auth {
	return &Auth{
		Username: username,
		Password: password,
	}
}

// GetEnvPassword gets password from OS environment
// It returns the password
func (m *Auth) GetEnvPassword() (string, error) {
	password := os.Getenv("FLEXVM_ACCESS_PASSWORD")

	if password == "" {
		return password, fmt.Errorf("GetEnvPassword error")
	}

	m.Password = password

	return password, nil
}

// GetEnvUsername gets FlexVM username from OS environment
// It returns the username
func (m *Auth) GetEnvUsername() (string, error) {
	h := os.Getenv("FLEXVM_ACCESS_USERNAME")

	if h == "" {
		return h, fmt.Errorf("GetEnvUsername error")
	}

	m.Username = h

	return h, nil
}
