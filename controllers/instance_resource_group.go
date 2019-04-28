package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/validation"
	"rs_web_go/common"
	"rs_web_go/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// InstanceResourceGroupController operations for InstanceResourceGroup
type InstanceResourceGroupController struct {
	beego.Controller
}

// URLMapping ...
func (c *InstanceResourceGroupController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetByPage)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create InstanceResourceGroup
// @Param	body		body 	models.InstanceResourceGroup	true		"body for InstanceResourceGroup content"
// @Success 201 {int} models.InstanceResourceGroup
// @Failure 403 body is empty
// @router / [post]
func (c *InstanceResourceGroupController) Post() {
	userIdStr := c.Ctx.GetCookie("userId")
	userId, _ := strconv.Atoi(userIdStr)
	var v models.InstanceResourceGroup
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.UserId = int64(userId)
		valid := validation.Validation{}
		b, err := valid.Valid(&v)
		if err != nil {
			c.Data["json"] = common.NewErrorResponse(err.Error())
			c.ServeJSON()
			return
		}
		if !b {
			errorStrs := make([]string, 0)
			for _, err := range valid.Errors {
				errorStrs = append(errorStrs, err.Message)
			}
			c.Data["json"] = common.NewErrorResponse(strings.Join(errorStrs, ","))
			c.ServeJSON()
			return
		}
		if _, err := models.AddInstanceResourceGroup(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.NewSuccessResponse(nil)
		} else {
			c.Data["json"] = common.NewErrorResponse(err.Error())
		}
	} else {
		c.Data["json"] = common.NewErrorResponse(err.Error())
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get InstanceResourceGroup by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.InstanceResourceGroup
// @Failure 403 :id is empty
// @router /:id [get]
func (c *InstanceResourceGroupController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetInstanceResourceGroupById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get InstanceResourceGroup
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.InstanceResourceGroup
// @Failure 403
// @router / [get]
func (c *InstanceResourceGroupController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllInstanceResourceGroup(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// GetByPage ...
// @Title Get By Page
// @Description get InstanceResourceGroup
// @Param	page	query	string	false	"Page number. Must be an integer"
// @Param	pageSize	query	string	false	"Limit the size of result set. Must be an integer"
// @Success 200 {object} models.InstanceResourceGroup
// @Failure 403
// @router / [get]
func (c *InstanceResourceGroupController) GetByPage() {
	userId := c.Ctx.GetCookie("userId")
	defer common.HandleCrash(func (s interface{})() {
	    c.Data["json"] = common.NewErrorResponse(s)
		c.ServeJSON()
	})
	var pageSize int64 = 10
	var page int64 = 1

	// limit: 10 (default is 10)
	if v, err := c.GetInt64("page"); err == nil {
		page = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("pageSize"); err == nil {
		pageSize = v
	}
	query := make(map[string]string)
	query["user_id"] = userId
	l, err := models.GetAllInstanceResourceGroupByPage(query,
		[]string{}, []string{"id"}, []string{"asc"}, page, pageSize)
	if err != nil {
		c.Data["json"] = common.NewErrorResponse(err.Error())
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the InstanceResourceGroup
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.InstanceResourceGroup	true		"body for InstanceResourceGroup content"
// @Success 200 {object} models.InstanceResourceGroup
// @Failure 403 :id is not int
// @router /:id [put]
func (c *InstanceResourceGroupController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.InstanceResourceGroup{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateInstanceResourceGroupById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the InstanceResourceGroup
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *InstanceResourceGroupController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteInstanceResourceGroup(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
