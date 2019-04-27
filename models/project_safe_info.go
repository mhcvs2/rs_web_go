package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ProjectSafeInfo struct {
	Id                     int       `orm:"column(id);auto"`
	ProjectId              int       `orm:"column(project_id);null"`
	DevConfigAlert         int8      `orm:"column(dev_config_alert);null" description:"允许开发人员配置报警 1 是 0 否"`
	DevCommitSchedulerTask int8      `orm:"column(dev_commit_scheduler_task);null" description:"允许开发人员提交任务调度  1 是 0 否"`
	DevExecuteTask         int8      `orm:"column(dev_execute_task);null" description:"允许开发人员立即执行任务   1 是 0 否"`
	WorkflowSaveVersionNum int       `orm:"column(workflow_save_version_num);null" description:"工作流保留版本数"`
	CreateTime             time.Time `orm:"column(create_time);type(timestamp);null;auto_now" description:"创建时间"`
	UpdateTime             time.Time `orm:"column(update_time);type(timestamp);null;auto_now" description:"更新时间"`
	IsDel                  int8      `orm:"column(is_del);null" description:"删除标识 0 未删除 1 已删除  默认为0"`
}

func (t *ProjectSafeInfo) TableName() string {
	return "project_safe_info"
}

func init() {
	orm.RegisterModel(new(ProjectSafeInfo))
}

// AddProjectSafeInfo insert a new ProjectSafeInfo into database and returns
// last inserted Id on success.
func AddProjectSafeInfo(m *ProjectSafeInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProjectSafeInfoById retrieves ProjectSafeInfo by Id. Returns error if
// Id doesn't exist
func GetProjectSafeInfoById(id int) (v *ProjectSafeInfo, err error) {
	o := orm.NewOrm()
	v = &ProjectSafeInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProjectSafeInfo retrieves all ProjectSafeInfo matches certain condition. Returns empty list if
// no records exist
func GetAllProjectSafeInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProjectSafeInfo))
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

	var l []ProjectSafeInfo
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

// UpdateProjectSafeInfo updates ProjectSafeInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectSafeInfoById(m *ProjectSafeInfo) (err error) {
	o := orm.NewOrm()
	v := ProjectSafeInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProjectSafeInfo deletes ProjectSafeInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProjectSafeInfo(id int) (err error) {
	o := orm.NewOrm()
	v := ProjectSafeInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProjectSafeInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
