package user_service

import (
	m "github.com/joaco-basile/pedidocs-back/models"
	user_repository "github.com/joaco-basile/pedidocs-back/repository/user.repositoy"
)

func Register(user m.User) error {

	err := user_repository.RegisterUser(user)

	if err != nil {
		return err
	}

	return nil
}

func Login(name, email, password string) (m.User, error) {

	user, err := user_repository.LoginUser(name, email, password)

	if err != nil {
		return user, err
	}

	return user, nil
}

func Update(user m.User, userId string) error {

	err := user_repository.UpdateUser(user, userId)

	if err != nil {
		return err
	}

	return nil
}

func Delete(userId string) error {

	err := user_repository.DseleteUser(userId)

	if err != nil {
		return err
	}

	return nil
}
