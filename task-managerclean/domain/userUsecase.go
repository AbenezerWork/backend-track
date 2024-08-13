package domain

type UserUsecase interface {
	AddUser(user User) (string, error)
	RemoveUser(sID string) error
	UserLogin(user *User) error
}
