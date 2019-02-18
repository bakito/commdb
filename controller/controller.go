package controller

import (
	"github.com/bakito/commdb/types"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

// New a new controller
func New(orm *xorm.Engine) *CommandController {
	return &CommandController{orm: orm}
}

// CommandController command controller
type CommandController struct {
	orm *xorm.Engine
}

// GetBy get command by ID
func (c *CommandController) GetBy(id int64) interface{} {
	command := &types.Command{ID: id}
	if ok, _ := c.orm.Get(command); ok {
		return command
	}

	return iris.StatusNotFound
}

// DeleteBy delete command by ID
func (c *CommandController) DeleteBy(id int64, ctx iris.Context) interface{} {
	command := &types.Command{}
	_, err := c.orm.ID(id).Delete(command)

	if err != nil {
		ctx.Application().Logger().Error(err)
		return iris.StatusInternalServerError
	}

	return nil

}

// Get get all existing commands
func (c *CommandController) Get() interface{} {
	commands := []types.Command{}
	c.orm.Find(&commands)
	return commands
}

// Put create a new command
func (c *CommandController) Put(ctx iris.Context) interface{} {

	command := &types.Command{}
	err := ctx.ReadJSON(command)
	_, err = c.orm.Insert(command)
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
		return iris.StatusInternalServerError
	}

	_, err = c.orm.ID(id).Update(command)
	if err == nil {
		return command
	}
	ctx.Application().Logger().Error(err)

	return iris.StatusInternalServerError
}
