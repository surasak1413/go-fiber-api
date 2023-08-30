package repository

type AuthModel struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	Password   string `json:"password"`
	Updatetime int64  `json:"update_time"`
	CreatedAt  int64  `json:"created_at"`
}
