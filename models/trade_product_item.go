package models

type TradeProductItem struct {
	ProductId   string `orm:"column(product_id);size(50);null"`
	ProductType int    `orm:"column(product_type);null"`
	UserId      int64  `orm:"column(user_id);null"`
	Region      string `orm:"column(region);size(50);null"`
	ItemNo      string `orm:"column(item_no);size(50);null"`
	ItemName    string `orm:"column(item_name);size(50);null"`
	Value       string `orm:"column(value);size(200);null"`
	ValueDev    string `orm:"column(value_dev);size(200);null"`
	ValueProd   string `orm:"column(value_prod);size(200);null"`
}
