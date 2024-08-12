package domain

type UserUsecase interface {
	InitUserUsecase(repository *UserRepository)
	AddUser(user User) (string, error)
	RemoveUser(sID string) error
	UserLogin(user *User) error
}
