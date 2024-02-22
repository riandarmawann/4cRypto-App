package config

const (
	UserSession      = "user"
	UserAdmin        = "id_user_dengan_role_admin"
	AuthGroup        = "/auth"
	AuthRegister     = "/register"
	AuthLogin        = "/login"
	AuthRefreshToken = "/refresh-token"

	UserGroup      = "/users"
	UpdateUserByID = "/update/:id"
	DeleteUserByID = "/delete/:id"
	CreateUser     = "/create"
	UserGetByID    = "/get/:id"
)
