package services

import (
	"loly/repository"
)

type service struct {
	rp repository.IAction
}

func NewService() *service {
	return &service{}
}
