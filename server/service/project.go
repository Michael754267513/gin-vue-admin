package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateProject
// @description   create a Project
// @param     project               model.Project
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateProject(project model.Project) (err error) {
	err = global.GVA_DB.Create(&project).Error
	return err
}

// @title    DeleteProject
// @description   delete a Project
// @auth                     （2020/04/05  20:22）
// @param     project               model.Project
// @return                    error

func DeleteProject(project model.Project) (err error) {
	err = global.GVA_DB.Delete(project).Error
	return err
}

// @title    DeleteProjectByIds
// @description   delete Projects
// @auth                     （2020/04/05  20:22）
// @param     project               model.Project
// @return                    error

func DeleteProjectByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Project{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateProject
// @description   update a Project
// @param     project          *model.Project
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateProject(project *model.Project) (err error) {
	err = global.GVA_DB.Save(project).Error
	return err
}

// @title    GetProject
// @description   get the info of a Project
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Project        Project

func GetProject(id uint) (err error, project model.Project) {
	err = global.GVA_DB.Where("id = ?", id).First(&project).Error
	return
}

// @title    GetProjectInfoList
// @description   get Project list by pagination, 分页获取Project
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetProjectInfoList(info request.ProjectSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Project{})
	var projects []model.Project
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&projects).Error
	return err, projects, total
}
