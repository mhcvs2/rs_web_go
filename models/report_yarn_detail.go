package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ReportYarnDetail struct {
	Id            int       `orm:"column(id);auto"`
	UserId        int64     `orm:"column(user_id)"`
	ProductTypeId int       `orm:"column(product_type_id)"`
	RegionCode    string    `orm:"column(region_code);size(50);null"`
	ItemNo        string    `orm:"column(item_no);size(50)"`
	ResourceConf  string    `orm:"column(resource_conf);size(255)"`
	MeasureDate   time.Time `orm:"column(measure_date);type(date);null"`
	MeasureTime   time.Time `orm:"column(measure_time);type(datetime)"`
	MeasureValue  string    `orm:"column(measure_value);size(255)"`
	Timestamp     time.Time `orm:"column(timestamp);type(timestamp);auto_now"`
}

func (t *ReportYarnDetail) TableName() string {
	return "report_yarn_detail"
}

func init() {
	orm.RegisterModel(new(ReportYarnDetail))
}

// AddReportYarnDetail insert a new ReportYarnDetail into database and returns
// last inserted Id on success.
func AddReportYarnDetail(m *ReportYarnDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReportYarnDetailById retrieves ReportYarnDetail by Id. Returns error if
// Id doesn't exist
func GetReportYarnDetailById(id int) (v *ReportYarnDetail, err error) {
	o := orm.NewOrm()
	v = &ReportYarnDetail{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReportYarnDetail retrieves all ReportYarnDetail matches certain condition. Returns empty list if
// no records exist
func GetAllReportYarnDetail(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ReportYarnDetail))
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

	var l []ReportYarnDetail
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

// UpdateReportYarnDetail updates ReportYarnDetail by Id and returns error if
// the record to be updated doesn't exist
func UpdateReportYarnDetailById(m *ReportYarnDetail) (err error) {
	o := orm.NewOrm()
	v := ReportYarnDetail{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReportYarnDetail deletes ReportYarnDetail by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReportYarnDetail(id int) (err error) {
	o := orm.NewOrm()
	v := ReportYarnDetail{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ReportYarnDetail{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
