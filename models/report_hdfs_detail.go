package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ReportHdfsDetail struct {
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

func (t *ReportHdfsDetail) TableName() string {
	return "report_hdfs_detail"
}

func init() {
	orm.RegisterModel(new(ReportHdfsDetail))
}

// AddReportHdfsDetail insert a new ReportHdfsDetail into database and returns
// last inserted Id on success.
func AddReportHdfsDetail(m *ReportHdfsDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReportHdfsDetailById retrieves ReportHdfsDetail by Id. Returns error if
// Id doesn't exist
func GetReportHdfsDetailById(id int) (v *ReportHdfsDetail, err error) {
	o := orm.NewOrm()
	v = &ReportHdfsDetail{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReportHdfsDetail retrieves all ReportHdfsDetail matches certain condition. Returns empty list if
// no records exist
func GetAllReportHdfsDetail(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ReportHdfsDetail))
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

	var l []ReportHdfsDetail
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

// UpdateReportHdfsDetail updates ReportHdfsDetail by Id and returns error if
// the record to be updated doesn't exist
func UpdateReportHdfsDetailById(m *ReportHdfsDetail) (err error) {
	o := orm.NewOrm()
	v := ReportHdfsDetail{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReportHdfsDetail deletes ReportHdfsDetail by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReportHdfsDetail(id int) (err error) {
	o := orm.NewOrm()
	v := ReportHdfsDetail{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ReportHdfsDetail{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
