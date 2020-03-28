package user

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"orderStatistics/request"
	overtimeService "orderStatistics/service/overtime"
	"orderStatistics/transform"
)

type OvertimeController struct{}

// @tags 加班
// @summary 吃加班餐登记
// @success 201
// @route /overtime/join [post]
func (c *OvertimeController) PostJoin(ctx iris.Context) error {
	var params request.JoinOvertime
	if err := ctx.ReadJSON(&params); err != nil {
		return err
	}
	if err := overtimeService.Service().JoinToday(params.UserID); err != nil {
		return err
	}
	ctx.StatusCode(http.StatusCreated)
	return nil
}

// @tags 加班
// @summary 今天登记列表
// @success 200
// @route /overtime/today [get]
func (c *OvertimeController) GetToday(ctx iris.Context) error {
	records, err := overtimeService.Service().GetTodayJoinedUsers()
	if err != nil {
		return err
	}

	return JSON(ctx, transform.User().ConvertUserList(records))
}
