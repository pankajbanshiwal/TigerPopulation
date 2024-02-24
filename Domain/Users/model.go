package Users

type CreateUser struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Token    string `json:"auth_token"`
	Status   string `json:"status"`
}
