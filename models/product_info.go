package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ProductInfo struct {
	Id          int    `orm:"column(id);auto"`
	Name        string `orm:"column(name);size(50)" description:"产品信息"`
	Code        string `orm:"column(code);size(50);null"`
	Description string `orm:"column(description);size(512);null" description:"描述信息"`
	LogoPath    string `orm:"column(logo_path);size(255);null" description:"logo路径"`
	DetailsPath string `orm:"column(details_path);size(255);null" description:"产品说明详情地址"`
	FlowType    int    `orm:"column(flow_type);null" description:"工作流 类型"`
}

func (t *ProductInfo) TableName() string {
	return "product_info"
}

func init() {
	orm.RegisterModel(new(ProductInfo))
}

// AddProductInfo insert a new ProductInfo into database and returns
// last inserted Id on success.
func AddProductInfo(m *ProductInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProductInfoById retrieves ProductInfo by Id. Returns error if
// Id doesn't exist
func GetProductInfoById(id int) (v *ProductInfo, err error) {
	o := orm.NewOrm()
	v = &ProductInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProductInfo retrieves all ProductInfo matches certain condition. Returns empty list if
// no records exist
func GetAllProductInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductInfo))
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

	var l []ProductInfo
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

// UpdateProductInfo updates ProductInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductInfoById(m *ProductInfo) (err error) {
	o := orm.NewOrm()
	v := ProductInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProductInfo deletes ProductInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProductInfo(id int) (err error) {
	o := orm.NewOrm()
	v := ProductInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
