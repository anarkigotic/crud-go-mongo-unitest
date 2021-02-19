package user_service

import (
	model "../../models"
	userRepository "../../repositories/user.repository"
)

func Create(user model.User) error {
	err := userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func Read() (model.Users, error) {
	users, err := userRepository.Read()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func Update(user model.User, userId string) error {
	err := userRepository.Update(user, userId)
	if err != nil {
		return nil
	}
	return nil
}

func Delete(userId string) error {
	err := userRepository.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
