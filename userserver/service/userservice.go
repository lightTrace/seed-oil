package service

import "errors"

type IUserService interface{
	GetName(userId int) string
	DelName(userId int) error
}

type UserService struct{}

func(us *UserService) GetName(userId int)string{
	if userId == 101{
		return "ming"
	}
	return "gang"
}

func(us *UserService) DelName(userId int)error{
	if userId == 101{
		return errors.New("无权限")
	}
	return nil
}
