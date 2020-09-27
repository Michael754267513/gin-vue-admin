// 自动生成模板Project
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Project struct {
	gorm.Model
	Name string `json:"NameID" form:"NameID" gorm:"column:name;comment:项目名称;type:varchar(256);size:256;"`
	Desc string `json:"Desc" form:"Desc" gorm:"column:desc;comment:描述;type:text(1024);size:1024;"`
}

func (Project) TableName() string {
	return "project"
}
