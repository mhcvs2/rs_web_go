package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ProjectInfo struct {
	Id          int       `orm:"column(id);auto"`
	Name        string    `orm:"column(name);size(100)" description:"项目描述"`
	Description string    `orm:"column(description);size(512);null" description:"项目描述"`
	OwnerId     int64     `orm:"column(owner_id);null" description:"项目创建人"`
	OwnerName   string    `orm:"column(owner_name);size(50);null" description:"项目创建者名称"`
	CreateTime  time.Time `orm:"column(create_time);type(datetime)" description:"创建时间"`
	UpdateTime  time.Time `orm:"column(update_time);type(timestamp);auto_now" description:"更新时间"`
	UpdateBy    string    `orm:"column(update_by);size(50);null" description:"更新者唯一标识符（主账号ID或者是子账号ID）"`
	MasterId    string    `orm:"column(master_id);size(50);null" description:"项目具体负责人的唯一标识符  默认项目创建人即为项目负责人"`
	MasterName  string    `orm:"column(master_name);size(50);null" description:"项目负责人名称"`
	IsDel       int8      `orm:"column(is_del);null" description:"删除标识 0 未删除   1 已删除 默认为0 "`
	UseStatus   int8      `orm:"column(use_status);null" description:"使用状态标识 0 禁用 1 启用 默认为 1"`
}

func (t *ProjectInfo) TableName() string {
	return "project_info"
}

func init() {
	orm.RegisterModel(new(ProjectInfo))
}

// AddProjectInfo insert a new ProjectInfo into database and returns
// last inserted Id on success.
func AddProjectInfo(m *ProjectInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProjectInfoById retrieves ProjectInfo by Id. Returns error if
// Id doesn't exist
func GetProjectInfoById(id int) (v *ProjectInfo, err error) {
	o := orm.NewOrm()
	v = &ProjectInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProjectInfo retrieves all ProjectInfo matches certain condition. Returns empty list if
// no records exist
func GetAllProjectInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProjectInfo))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ProjectInfo
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateProjectInfo updates ProjectInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectInfoById(m *ProjectInfo) (err error) {
	o := orm.NewOrm()
	v := ProjectInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProjectInfo deletes ProjectInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProjectInfo(id int) (err error) {
	o := orm.NewOrm()
	v := ProjectInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProjectInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
