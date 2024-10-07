package auth

import (
	"fmt"
	"os"
)

// Auth describes the authentication information for FortiFlex
type Auth struct {
	Username string
	Password string
	Token    string
}

// NewAuth inits Auth object with the given metadata
func NewAuth(username, password string) (*Auth, error) {
	auth := &Auth{
		Username: username,
		Password: password,
	}
	if auth.Username == "" {
		_, err := auth.getEnvUsername()
		if err != nil {
			return nil, fmt.Errorf("Error reading Username")
		}
	}
	if auth.Password == "" {
		_, err := auth.getEnvPassword()
		if err != nil {
			return nil, fmt.Errorf("Error reading Password")
		}
	}
	return auth, nil
}

// getEnvPassword gets password from OS environment
// It returns the password
func (m *Auth) getEnvPassword() (string, error) {
	password := os.Getenv("FORTIFLEX_ACCESS_PASSWORD")

	if password == "" {
		return password, fmt.Errorf("GetEnvPassword error")
	}

	m.Password = password

	return password, nil
}

// getEnvUsername gets username from OS environment
// It returns the username
func (m *Auth) getEnvUsername() (string, error) {
	h := os.Getenv("FORTIFLEX_ACCESS_USERNAME")

	if h == "" {
		return h, fmt.Errorf("GetEnvUsername error")
	}

	m.Username = h

	return h, nil
}
