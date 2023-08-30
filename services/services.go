package services

import (
	"loly/database"
	"loly/repository"
)

type Action interface {
	Validate() error
	AuthService() AuthServiceAction
}

func (sv *service) Validate() error {

	return nil
}

type service struct {
	rp repository.IAction
}

func NewService() (Action, error) {
	sv := new(service)

	if sv.rp == nil {
		db, err := database.Connect()
		if err != nil {
			return nil, err
		}
		repo := repository.NewStore(db)
		sv.rp = repo
	}

	return sv, nil
}
