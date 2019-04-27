package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TradeInstanceResource struct {
	Id              int    `orm:"column(instance_resource_id);auto"`
	InstanceId      string `orm:"column(instance_id);size(255);null"`
	ItemNo          string `orm:"column(item_No);size(50);null"`
	Value           string `orm:"column(value);size(50);null"`
	ResourceConf    string `orm:"column(resource_conf);size(255);null" description:"实例计费项对应的资源配置"`
	IsResourceGroup int    `orm:"column(is_resource_group);null" description:"是否有资源组，0 无 1 有"`
	ValueDev        string `orm:"column(value_dev);size(200);null"`
	ValueProd       string `orm:"column(value_prod);size(200);null"`
}

func (t *TradeInstanceResource) TableName() string {
	return "trade_instance_resource"
}

func init() {
	orm.RegisterModel(new(TradeInstanceResource))
}

// AddTradeInstanceResource insert a new TradeInstanceResource into database and returns
// last inserted Id on success.
func AddTradeInstanceResource(m *TradeInstanceResource) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTradeInstanceResourceById retrieves TradeInstanceResource by Id. Returns error if
// Id doesn't exist
func GetTradeInstanceResourceById(id int) (v *TradeInstanceResource, err error) {
	o := orm.NewOrm()
	v = &TradeInstanceResource{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTradeInstanceResource retrieves all TradeInstanceResource matches certain condition. Returns empty list if
// no records exist
func GetAllTradeInstanceResource(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TradeInstanceResource))
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

	var l []TradeInstanceResource
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

// UpdateTradeInstanceResource updates TradeInstanceResource by Id and returns error if
// the record to be updated doesn't exist
func UpdateTradeInstanceResourceById(m *TradeInstanceResource) (err error) {
	o := orm.NewOrm()
	v := TradeInstanceResource{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTradeInstanceResource deletes TradeInstanceResource by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTradeInstanceResource(id int) (err error) {
	o := orm.NewOrm()
	v := TradeInstanceResource{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TradeInstanceResource{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
