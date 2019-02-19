package controller

import (
	"strconv"

	"github.com/bakito/commdb/service"
	"github.com/bakito/commdb/types"
	"github.com/kataras/iris"
)

// New a new controller
func New() *CommandController {
	return &CommandController{}
}

// CommandController command controller
type CommandController struct {
	Service service.CommandService
}

// GetBy get command by ID
func (c *CommandController) GetBy(id int64) interface{} {
	if command, ok := c.Service.GetByID(id); ok {
		return command
	}

	return iris.StatusNotFound
}

// DeleteBy delete command by ID
func (c *CommandController) DeleteBy(id int64, ctx iris.Context) interface{} {
	err := c.Service.DeleteByID(id)

	if err != nil {
		ctx.Application().Logger().Error(err)
		return iris.StatusInternalServerError
	}

	return nil
}

// Get get all existing commands
func (c *CommandController) Get(ctx iris.Context) interface{} {
	pageSize, err := strconv.Atoi(ctx.URLParamDefault("pageSize", "100"))

	if err != nil {
		ctx.Application().Logger().Error(err)
		return iris.StatusBadRequest
	}

	page, err := strconv.Atoi(ctx.URLParamDefault("page", "0"))

	if err != nil {
		ctx.Application().Logger().Error(err)
		return iris.StatusBadRequest
	}

	query := ctx.URLParam("search")

	commands, err := c.Service.GetAll(query, page, pageSize)

	if err != nil {
		ctx.Application().Logger().Error(err)
		return iris.StatusInternalServerError
	}

	return commands
}

// Put create a new command
func (c *CommandController) Put(ctx iris.Context) interface{} {

	command := &types.Command{}
	err := ctx.ReadJSON(command)

	if err != nil {
		ctx.Application().Logger().Error(err)
		return iris.StatusBadRequest
	}

	_, err = c.Service.Create(command)

	if err == nil {
		return command
	}

	ctx.Application().Logger().Error(err)

	return iris.StatusInternalServerError
}

// PostBy update a command
func (c *CommandController) PostBy(id int64, ctx iris.Context) interface{} {

	command := &types.Command{}
	err := ctx.ReadJSON(command)
	if err != nil {
		ctx.Application().Logger().Error(err)
		return iris.StatusBadRequest
	}

	command.ID = id
	err = c.Service.Update(command)
	if err == nil {
		return command
	}
	ctx.Application().Logger().Error(err)

	return iris.StatusInternalServerError
}
