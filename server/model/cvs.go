// 自动生成模板Cvs
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Cvs struct {
	gorm.Model
	CVSname      string `json:"CVSname" form:"CVSname" gorm:"column:CVSname;comment:版本控制器;type:varchar(256);size:256;"`
	CVSCachePath string `json:"CVSCachePath" form:"CVSCachePath" gorm:"column:CVSCachePath;comment:缓存路径;type:varchar(256);size:256;"`
	CVSBuildPath string `json:"CVSBuildPath" form:"CVSBuildPath" gorm:"column:CVSBuildPath;comment:编译路径;type:varchar(256);size:256;"`
}

func (Cvs) TableName() string {
	return "csv"
}
