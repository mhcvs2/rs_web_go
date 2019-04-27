package models

import "time"

type DiStoCount struct {
	TenantId    int64     `orm:"column(tenant_id)"`
	ProjectId   int       `orm:"column(project_id)"`
	TopicName   string    `orm:"column(topic_name);size(50)"`
	MeasureTime time.Time `orm:"column(measure_time);type(datetime)"`
	StorageSize int64     `orm:"column(storage_size)"`
}
