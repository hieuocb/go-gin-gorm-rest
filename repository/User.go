package repository

type IUserRepository interface {
	GetAllUser()
	GetUserById(id string)
	UpdateUserById(id string)
	DeleteUserById(id string)
}
