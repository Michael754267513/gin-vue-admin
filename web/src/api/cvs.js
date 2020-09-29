import service from '@/utils/request'

// @Tags Cvs
// @Summary 创建Cvs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cvs true "创建Cvs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csv/createCvs [post]
export const createCvs = (data) => {
     return service({
         url: "/csv/createCvs",
         method: 'post',
         data
     })
 }


// @Tags Cvs
// @Summary 删除Cvs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cvs true "删除Cvs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csv/deleteCvs [delete]
 export const deleteCvs = (data) => {
     return service({
         url: "/csv/deleteCvs",
         method: 'delete',
         data
     })
 }

// @Tags Cvs
// @Summary 删除Cvs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cvs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csv/deleteCvs [delete]
 export const deleteCvsByIds = (data) => {
     return service({
         url: "/csv/deleteCvsByIds",
         method: 'delete',
         data
     })
 }

// @Tags Cvs
// @Summary 更新Cvs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cvs true "更新Cvs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csv/updateCvs [put]
 export const updateCvs = (data) => {
     return service({
         url: "/csv/updateCvs",
         method: 'put',
         data
     })
 }


// @Tags Cvs
// @Summary 用id查询Cvs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cvs true "用id查询Cvs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csv/findCvs [get]
 export const findCvs = (params) => {
     return service({
         url: "/csv/findCvs",
         method: 'get',
         params
     })
 }


// @Tags Cvs
// @Summary 分页获取Cvs列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Cvs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csv/getCvsList [get]
 export const getCvsList = (params) => {
     return service({
         url: "/csv/getCvsList",
         method: 'get',
         params
     })
 }