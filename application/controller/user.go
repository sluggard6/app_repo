package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/shfd/repo/model"
	"github.com/shfd/repo/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController() *UserController {
	return &UserController{service.NewUserService()}
}

type LoginForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResetPassForm struct {
	Userid  uint   `json:"userid"`
	Oldpass string `json:"oldpass"`
	Newpass string `json:"newpass" validate:"required"`
}

type LoginInfo struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}

func (c *UserController) PostLogin(ctx iris.Context) *HttpResult {
	loginForm := &LoginForm{}
	if err := ctx.ReadJSON(loginForm); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return FailedCode(PARAM_ERROR)
	}
	if user, err := c.userService.Login(loginForm.Username, loginForm.Password); user == nil {
		return FailedCodeMessage(LOGIN_FAILED, err.Error())
	} else {
		session := sessions.Get(ctx)
		session.Set("authenticated", true)
		session.Set("user", user)
		return Success(&LoginInfo{User: user, Token: session.ID()})
	}
}

func (c *UserController) PostLogout(ctx iris.Context) *HttpResult {
	session := sessions.Get(ctx)
	session.Destroy()
	return Success(nil)
}

func (c *UserController) GetInfo(ctx iris.Context) *HttpResult {
	session := sessions.Get(ctx)
	return c.GetBy(session.Get("user").(*model.User).ID)
}

func (c *UserController) GetBy(id uint) *HttpResult {
	if user, err := c.userService.GetById(id); err != nil {
		return Failed()
	} else {
		user.Password = ""
		return Success(user)
	}
}

func (c *UserController) PostRegister(ctx iris.Context) *HttpResult {
	user := &model.User{}
	if err := ctx.ReadJSON(user); err != nil {
		return FailedCode(PARAM_ERROR)
	}
	if err := validate.Struct(user); err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return FailedCodeMessage(PARAM_ERROR, err.Error())
		}

		// from here you can create your own error messages in whatever language you wish
		return FailedCodeMessage(PARAM_ERROR, err.Error())
	}
	if err := c.userService.Register(user); err != nil {
		return FailedCodeMessage(LOGIN_FAILED, err.Error())
	}
	return Success(user)
}

func (c *UserController) GetLike(ctx iris.Context) *HttpResult {
	user := sessions.Get(ctx).Get("user").(*model.User)
	queryString := ctx.URLParam("queryString")
	s := make([]model.User, 0)
	result := &s
	model.DB().Table("users").Select("id", "username").Where("username like ? AND id <> ? AND `users`.`deleted_at` IS NULL", "%"+queryString+"%", user.ID).Scan(result)
	// model.DB().Where("username like ? AND id <> ?", "%"+queryString+"%", user.ID).Find(result)
	return Success(result)
}

func (c *UserController) PostResetpass(ctx iris.Context) *HttpResult {
	user := sessions.Get(ctx).Get("user").(*model.User)
	resetPassForm := &ResetPassForm{}
	if err := ctx.ReadJSON(resetPassForm); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return FailedCode(PARAM_ERROR)
	}
	res, err := c.userService.ResetPassword(user.ID, resetPassForm.Oldpass, resetPassForm.Newpass)
	if res {
		return Success("success")
	} else {
		if err == nil {
			return FailedMessage("password wrong")
		} else {
			return FailedMessage(err.Error())
		}

	}
}

func (c *UserController) GetTest(ctx iris.Context) *HttpResult {
	return Success("test success")
}
