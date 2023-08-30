package repository

import (
	"fmt"
)

type authRepository interface {
	RegisterAction(inp *AuthModel) error
}

func (st *store) RegisterAction(inp *AuthModel) error {
	fmt.Println("RegisterAction")
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inp.Password), 14)
	// if err != nil {
	// 	return err
	// }

	// st.db.Create(&AuthModel{
	// 	Id:         inp.Id,
	// 	Username:   inp.Username,
	// 	Password:   string(hashedPassword),
	// 	Role:       inp.Role,
	// 	Updatetime: inp.Updatetime,
	// 	CreatedAt:  inp.CreatedAt,
	// })
	// st.db.Create(inp)

	return nil
}
