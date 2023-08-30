package services

import (
	"fmt"
	"loly/repository"
	"loly/request"
	"loly/utils"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(req request.AuthRequest) error
}

func (sv *service) Register(req request.AuthRequest) error {
	fmt.Println("Register Service")
	currentTime := utils.GetCurrentEpochTime()
	insertData := &repository.AuthModel{
		Id:         uuid.NewString(),
		Username:   req.Username,
		Password:   req.Password,
		Role:       "USER",
		Updatetime: currentTime,
		CreatedAt:  currentTime,
	}
	fmt.Println(insertData)
	sv.rp.RegisterAction(insertData)

	return nil
}
