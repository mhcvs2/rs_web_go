package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type UserProductOrderInfo struct {
	Id            int       `orm:"column(id);auto" description:"订单id"`
	UserId        int64     `orm:"column(user_id)" description:"用户id"`
	ProductTypeId int       `orm:"column(product_type_id)" description:"产品类型di"`
	BillType      int       `orm:"column(bill_type);null" description:"计费方式"`
	RegionCode    string    `orm:"column(region_code);size(255);null" description:"机房code"`
	ItemNoList    string    `orm:"column(item_no_list);null" description:"配置明细 json格式"`
	OrderStatus   int8      `orm:"column(order_status);null" description:"订单状态  0未开通 1 开通"`
	UpdateTime    time.Time `orm:"column(update_time);type(timestamp);null;auto_now" description:"更新时间"`
	CreateTime    time.Time `orm:"column(create_time);type(datetime);null" description:"订单创建时间"`
	ProductId     string    `orm:"column(product_id);size(50);null" description:"商品id"`
}

func (t *UserProductOrderInfo) TableName() string {
	return "user_product_order_info"
}

func init() {
	orm.RegisterModel(new(UserProductOrderInfo))
}

// AddUserProductOrderInfo insert a new UserProductOrderInfo into database and returns
// last inserted Id on success.
func AddUserProductOrderInfo(m *UserProductOrderInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserProductOrderInfoById retrieves UserProductOrderInfo by Id. Returns error if
// Id doesn't exist
func GetUserProductOrderInfoById(id int) (v *UserProductOrderInfo, err error) {
	o := orm.NewOrm()
	v = &UserProductOrderInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUserProductOrderInfo retrieves all UserProductOrderInfo matches certain condition. Returns empty list if
// no records exist
func GetAllUserProductOrderInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserProductOrderInfo))
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

	var l []UserProductOrderInfo
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

// UpdateUserProductOrderInfo updates UserProductOrderInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserProductOrderInfoById(m *UserProductOrderInfo) (err error) {
	o := orm.NewOrm()
	v := UserProductOrderInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUserProductOrderInfo deletes UserProductOrderInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserProductOrderInfo(id int) (err error) {
	o := orm.NewOrm()
	v := UserProductOrderInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UserProductOrderInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
