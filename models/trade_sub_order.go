package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TradeSubOrder struct {
	Id           int    `orm:"column(sub_order_id);pk"`
	OrderId      string `orm:"column(order_id);size(50)"`
	ProductType  string `orm:"column(product_type);size(255)"`
	ProductGroup string `orm:"column(product_group);size(255)"`
	Region       string `orm:"column(region);size(255)"`
	ProductId    string `orm:"column(product_id);size(50)"`
	InstanceId   string `orm:"column(instance_id);size(50)" description:"实例id"`
	UserId       int64  `orm:"column(user_id)" description:"用户id"`
}

func (t *TradeSubOrder) TableName() string {
	return "trade_sub_order"
}

func init() {
	orm.RegisterModel(new(TradeSubOrder))
}

// AddTradeSubOrder insert a new TradeSubOrder into database and returns
// last inserted Id on success.
func AddTradeSubOrder(m *TradeSubOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTradeSubOrderById retrieves TradeSubOrder by Id. Returns error if
// Id doesn't exist
func GetTradeSubOrderById(id int) (v *TradeSubOrder, err error) {
	o := orm.NewOrm()
	v = &TradeSubOrder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTradeSubOrder retrieves all TradeSubOrder matches certain condition. Returns empty list if
// no records exist
func GetAllTradeSubOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TradeSubOrder))
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

	var l []TradeSubOrder
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

// UpdateTradeSubOrder updates TradeSubOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateTradeSubOrderById(m *TradeSubOrder) (err error) {
	o := orm.NewOrm()
	v := TradeSubOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTradeSubOrder deletes TradeSubOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTradeSubOrder(id int) (err error) {
	o := orm.NewOrm()
	v := TradeSubOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TradeSubOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
