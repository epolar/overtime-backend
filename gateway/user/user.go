package user

import (
	"github.com/kataras/iris/v12"
	"orderStatistics/request"
	userService "orderStatistics/service/user"
	"orderStatistics/transform"
)

type PersonController struct{}

// @tags 用户
// @summary 用户列表
// @success 200
// @route /user/all [get]
func (p *PersonController) GetAll(ctx iris.Context) error {
	list, err := userService.Service().All()
	if err != nil {
		return err
	}
	return JSON(ctx, transform.User().ConvertUserList(list))
}

// @tags 用户
// @summary 添加用户
// @success 201
// @route /user/add [post]
func (p *PersonController) PostAdd(ctx iris.Context) error {
	var params request.AddUser
	if err := ctx.ReadJSON(&params); err != nil {
		return err
	}
	if resp, err := userService.Service().Add(params.Name, params.Label); err != nil {
		return err
	} else {
		return JSON(ctx, transform.User().ConvertUser(resp))
	}
}
