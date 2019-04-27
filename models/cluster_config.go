package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ClusterConfig struct {
	Id            int       `orm:"column(id);auto"`
	ClusterName   string    `orm:"column(cluster_name);size(100);null" description:"集群名称"`
	Host          string    `orm:"column(host);size(100);null"`
	RootCuNum     int       `orm:"column(root_cu_num);null" description:"集群总cu数"`
	BasicKey      string    `orm:"column(basic_key);null"`
	Cookie        string    `orm:"column(cookie);size(255);null" description:"api请求需要携带的cookie信息"`
	RmHost        string    `orm:"column(rm_host);size(255);null" description:"resourceManager 地址"`
	NmHost        string    `orm:"column(nm_host);size(255);null"`
	ProductTypeId string    `orm:"column(product_type_id);size(255);null" description:"支持的产品类型"`
	UserId        string    `orm:"column(user_id);size(255);null" description:"支持的用户id"`
	IsDefault     int       `orm:"column(is_default);null" description:"1 是默认 0 否"`
	Status        int       `orm:"column(status);null" description:"状态 1 可以 0 禁用"`
	Timestamp     time.Time `orm:"column(timestamp);type(timestamp);null"`
	FileType      string    `orm:"column(file_type);size(255);null"`
	ClusterType   int       `orm:"column(cluster_type);null"`
	ClusterKind   int       `orm:"column(cluster_kind);null"`
}

func (t *ClusterConfig) TableName() string {
	return "cluster_config"
}

func init() {
	orm.RegisterModel(new(ClusterConfig))
}

// AddClusterConfig insert a new ClusterConfig into database and returns
// last inserted Id on success.
func AddClusterConfig(m *ClusterConfig) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetClusterConfigById retrieves ClusterConfig by Id. Returns error if
// Id doesn't exist
func GetClusterConfigById(id int) (v *ClusterConfig, err error) {
	o := orm.NewOrm()
	v = &ClusterConfig{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllClusterConfig retrieves all ClusterConfig matches certain condition. Returns empty list if
// no records exist
func GetAllClusterConfig(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ClusterConfig))
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

	var l []ClusterConfig
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

// UpdateClusterConfig updates ClusterConfig by Id and returns error if
// the record to be updated doesn't exist
func UpdateClusterConfigById(m *ClusterConfig) (err error) {
	o := orm.NewOrm()
	v := ClusterConfig{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteClusterConfig deletes ClusterConfig by Id and returns error if
// the record to be deleted doesn't exist
func DeleteClusterConfig(id int) (err error) {
	o := orm.NewOrm()
	v := ClusterConfig{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ClusterConfig{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
