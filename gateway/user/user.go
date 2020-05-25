package user

import (
	"orderStatistics/request"
	userService "orderStatistics/service/user"
	"orderStatistics/transform"

	"github.com/kataras/iris/v12"
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
	if resp, err := userService.Service().Add(params.Name, params.Label, params.Nick); err != nil {
		return err
	} else {
		return JSON(ctx, transform.User().ConvertUser(resp))
	}
}

// @tags 用户
// @summary 删除用户
// @param user_id path number true "用户ID"
// @success 204
// @route /user/{user_id} [delete]
func (p *PersonController) DeleteBy(userID uint64) error {
	return userService.Service().Delete(userID)
}
