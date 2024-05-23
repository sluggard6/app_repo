package controller

import "github.com/kataras/iris/v12"

type AppController struct {
}

func NewAppController() *AppController {
	return &AppController{}
}

func (c *AppController) PostCreateApp(ctx iris.Context) *HttpResult {
	return nil
}
