package service

import "github.com/shfd/repo/model"

type AppService interface {
	CreateApp(user *model.User) (*model.App, error)
}
