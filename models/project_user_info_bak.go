package models

type ProjectUserInfoBak struct {
	Id_RENAME int `orm:"column(id);null"`
	ProjectId int `orm:"column(project_id);null"`
	UserId    int `orm:"column(user_id);null"`
}
