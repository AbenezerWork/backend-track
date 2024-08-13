package usecase

import (
	"fmt"
	"task_managerv2/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepository domain.UserRepository
}

func InitUserUsecase(repository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{UserRepository: repository}
}

func (uu *UserUsecase) AddUser(user domain.User) (string, error) {
	id, err := uu.UserRepository.AddUser(&user)
	if err != nil {
		return "", err
	}
	return id.Hex(), nil
}

func (uu *UserUsecase) RemoveUser(sID string) error {
	id, err := primitive.ObjectIDFromHex(sID)
	if err != nil {
		return err
	}
	return uu.UserRepository.DeleteUser(id)
}

func (uu *UserUsecase) UserLogin(user *domain.User) error {
	userByEmail, err := uu.UserRepository.GetUserByEmail(user.Email)
	if err != nil {
		fmt.Println("userByEmail")
		return err
	}
	fmt.Println(bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost))
	fmt.Println([]byte(userByEmail.Password))

	err = bcrypt.CompareHashAndPassword([]byte(userByEmail.Password), []byte(user.Password))
	fmt.Println(err, "hello")
	if err != nil {
		fmt.Println("CompareHashAndPassword")
		return err
	}
	user.ID = userByEmail.ID
	user.Role = userByEmail.Role
	return nil
}
