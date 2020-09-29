package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags Cvs
// @Summary 创建Cvs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cvs true "创建Cvs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csv/createCvs [post]
func CreateCvs(c *gin.Context) {
	var csv model.Cvs
	_ = c.ShouldBindJSON(&csv)
	err := service.CreateCvs(csv)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Cvs
// @Summary 删除Cvs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cvs true "删除Cvs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csv/deleteCvs [delete]
func DeleteCvs(c *gin.Context) {
	var csv model.Cvs
	_ = c.ShouldBindJSON(&csv)
	err := service.DeleteCvs(csv)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Cvs
// @Summary 批量删除Cvs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cvs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csv/deleteCvsByIds [delete]
func DeleteCvsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteCvsByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Cvs
// @Summary 更新Cvs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cvs true "更新Cvs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csv/updateCvs [put]
func UpdateCvs(c *gin.Context) {
	var csv model.Cvs
	_ = c.ShouldBindJSON(&csv)
	err := service.UpdateCvs(&csv)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Cvs
// @Summary 用id查询Cvs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cvs true "用id查询Cvs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csv/findCvs [get]
func FindCvs(c *gin.Context) {
	var csv model.Cvs
	_ = c.ShouldBindQuery(&csv)
	err, recsv := service.GetCvs(csv.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"recsv": recsv}, c)
	}
}

// @Tags Cvs
// @Summary 分页获取Cvs列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CvsSearch true "分页获取Cvs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csv/getCvsList [get]
func GetCvsList(c *gin.Context) {
	var pageInfo request.CvsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetCvsInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
