package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
	"reflect"
	"rs_web_go/common"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type InstanceResourceGroup struct {
	Id            int       `orm:"column(id);auto" json:"id"`
	UserId        int64     `orm:"column(user_id)" json:"user_id" valid:"Required"`
	RegionCode    string    `orm:"column(region_code);size(50)" description:"机房code" json:"region_code" valid:"Required"`
	ProductTypeId int       `orm:"column(product_type_id)" description:"产品类型id" json:"product_type_id" valid:"Required"`
	GpName        string    `orm:"column(gp_name);size(50)" json:"gp_name" valid:"Required"`
	ItemNo        string    `orm:"column(item_no);size(255);null" description:"计费项" json:"item_no" valid:"Required"`
	ValueMax      int       `orm:"column(value_max);null" json:"value_max"`
	ValueMin      int       `orm:"column(value_min);null" json:"value_min"`
	Description   string    `orm:"column(description);size(255);null" json:"description"`
	Status        int16     `orm:"column(status);null" description:"分组状态 0删除 1 正常" json:"status"`
	CreateTime    time.Time `orm:"column(create_time);type(datetime);null;auto_now_add" json:"create_time"`
	UpdateTime    time.Time `orm:"column(update_time);type(timestamp);null;auto_now" json:"update_time"`
	ResourceConf  string    `orm:"column(resource_conf);size(255);null" description:"资源队列,真正的队列是id+resource_conf" json:"resource_conf"`
	IsDefault     int       `orm:"column(is_default);null" description:"是否默认资源分组 0 否 1是" json:"is_default"`
	ValueDev      int       `orm:"column(value_dev);null" json:"value_dev" valid:"Required; Min(0)"`
	ValueProd     int       `orm:"column(value_prod);null" json:"value_prod" valid:"Required; Min(0)"`
}

func (t *InstanceResourceGroup) TableName() string {
	return "instance_resource_group"
}

func (t *InstanceResourceGroup) Valid(v *validation.Validation) {
	if t.ItemNo != "CU" && t.ItemNo != "DCU" {
		v.SetError("ItemNo", "must be in [CU, DCU]")
	}
}

func init() {
	orm.RegisterModel(new(InstanceResourceGroup))
}

// AddInstanceResourceGroup insert a new InstanceResourceGroup into database and returns
// last inserted Id on success.
func AddInstanceResourceGroup(m *InstanceResourceGroup) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInstanceResourceGroupById retrieves InstanceResourceGroup by Id. Returns error if
// Id doesn't exist
func GetInstanceResourceGroupById(id int) (v *InstanceResourceGroup, err error) {
	o := orm.NewOrm()
	v = &InstanceResourceGroup{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInstanceResourceGroup retrieves all InstanceResourceGroup matches certain condition. Returns empty list if
// no records exist
func GetAllInstanceResourceGroup(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(InstanceResourceGroup))
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

	var l []InstanceResourceGroup
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

// GetAllInstanceResourceGroup retrieves all InstanceResourceGroup matches certain condition. Returns empty list if
// no records exist
func GetAllInstanceResourceGroupByPage(query map[string]string, fields []string, sortby []string, order []string,
	page int64, pageSize int64) (response *common.Response, err error) {
	ml := make([]interface{}, 0)
	o := orm.NewOrm()
	qs := o.QueryTable(new(InstanceResourceGroup))
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
	totalCount, err := qs.Count()
	if err != nil {
		return nil, err
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
	pageParam := common.NewPageParam(page, totalCount, pageSize)

	var l []InstanceResourceGroup
	qs = qs.OrderBy(sortFields...)
	offset := (page - 1) * pageSize
	if _, err = qs.Limit(pageSize, offset).All(&l, fields...); err == nil {
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
		return common.NewPageResponse(ml, "groups", pageParam), nil
	}
	return nil, err
}



// UpdateInstanceResourceGroup updates InstanceResourceGroup by Id and returns error if
// the record to be updated doesn't exist
func UpdateInstanceResourceGroupById(m *InstanceResourceGroup) (err error) {
	o := orm.NewOrm()
	v := InstanceResourceGroup{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInstanceResourceGroup deletes InstanceResourceGroup by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInstanceResourceGroup(id int) (err error) {
	o := orm.NewOrm()
	v := InstanceResourceGroup{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&InstanceResourceGroup{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
