package models

import "time"

type TradeInstanceInfo struct {
	Id_RENAME        int64     `orm:"column(id)"`
	UPDATETIME       time.Time `orm:"column(UPDATE_TIME);type(timestamp);null;auto_now"`
	CREATETIME       time.Time `orm:"column(CREATE_TIME);type(timestamp);null;auto_now_add"`
	INSTANCEID       string    `orm:"column(INSTANCE_ID);size(64)"`
	USERID           int64     `orm:"column(USER_ID);null"`
	REGION           string    `orm:"column(REGION);size(64);null"`
	PRODUCTTYPE      int       `orm:"column(PRODUCT_TYPE);null"`
	PRODUCTGROUP     int       `orm:"column(PRODUCT_GROUP);null"`
	BILLTYPE         int       `orm:"column(BILL_TYPE);null"`
	INSTANCETYPE     int       `orm:"column(INSTANCE_TYPE);null"`
	STATUS           int       `orm:"column(STATUS);null"`
	SERVICEBEGINTIME time.Time `orm:"column(SERVICE_BEGIN_TIME);type(timestamp);null"`
	SERVICEENDTIME   time.Time `orm:"column(SERVICE_END_TIME);type(datetime);null"`
}
