package request

import (
	"errors"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *AuthRequest) Validate() error {

	if req.Username == "" {
		return errors.New("username is required")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
