package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateCvs
// @description   create a Cvs
// @param     csv               model.Cvs
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateCvs(csv model.Cvs) (err error) {
	err = global.GVA_DB.Create(&csv).Error
	return err
}

// @title    DeleteCvs
// @description   delete a Cvs
// @auth                     （2020/04/05  20:22）
// @param     csv               model.Cvs
// @return                    error

func DeleteCvs(csv model.Cvs) (err error) {
	err = global.GVA_DB.Delete(csv).Error
	return err
}

// @title    DeleteCvsByIds
// @description   delete Cvss
// @auth                     （2020/04/05  20:22）
// @param     csv               model.Cvs
// @return                    error

func DeleteCvsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Cvs{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateCvs
// @description   update a Cvs
// @param     csv          *model.Cvs
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateCvs(csv *model.Cvs) (err error) {
	err = global.GVA_DB.Save(csv).Error
	return err
}

// @title    GetCvs
// @description   get the info of a Cvs
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Cvs        Cvs

func GetCvs(id uint) (err error, csv model.Cvs) {
	err = global.GVA_DB.Where("id = ?", id).First(&csv).Error
	return
}

// @title    GetCvsInfoList
// @description   get Cvs list by pagination, 分页获取Cvs
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetCvsInfoList(info request.CvsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Cvs{})
	var csvs []model.Cvs
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&csvs).Error
	return err, csvs, total
}
