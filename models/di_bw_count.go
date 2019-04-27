package models

import "time"

type DiBwCount struct {
	TenantId    int64     `orm:"column(tenant_id)"`
	ProjectId   int       `orm:"column(project_id)"`
	TopicName   string    `orm:"column(topic_name);size(50)"`
	MeasureTime time.Time `orm:"column(measure_time);type(datetime)"`
	Bandwidth   int64     `orm:"column(bandwidth)"`
}
