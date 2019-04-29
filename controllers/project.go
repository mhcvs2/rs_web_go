package controllers

import (
	"github.com/astaxie/beego"
	"rs_web_go/common"
	"rs_web_go/service"
)

// ProjectController operations for Project
type ProjectController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProjectController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetByPage ...
// @Title Get By Page
// @Description get Project
// @Success 200 {object}
// @Failure 403
// @router / [get]
func (c *ProjectController) GetAll() {
	defer common.HandleCrash(func (s interface{})() {
		c.Data["json"] = common.NewErrorResponse(s)
		c.ServeJSON()
	})

	rep, err := service.GetAllProjects()
	if err != nil {
		c.Data["json"] = common.NewErrorResponse(err.Error())
	} else {
		c.Data["json"] = common.NewSuccessResponse(rep)
	}
	c.ServeJSON()
}
